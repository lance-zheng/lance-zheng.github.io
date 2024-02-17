<!-- customize-category:Spring Framework -->

# Spring Security with JWT

- [Spring Security with JWT](#spring-security-with-jwt)
  - [Token 工具类封装](#token-工具类封装)
  - [Handler](#handler)
    - [AuthenticationSuccessHandler](#authenticationsuccesshandler)
    - [AuthenticationFailureHandler](#authenticationfailurehandler)
    - [AccessDeniedHandler](#accessdeniedhandler)
    - [AuthenticationEntryPoint](#authenticationentrypoint)
  - [UserDetailsService](#userdetailsservice)
  - [AuthenticationFilter](#authenticationfilter)
  - [最后的总配置类](#最后的总配置类)

[GitHub](https://github.com/lance-zheng/demo/tree/main/spring-security-jwt)

在 **Spring Security** 中使用 **JWT** 进行认证授权

## Token 工具类封装

Token 生成工具类，用来生成和解析 JWT Token

JWT 的实现有很多，具体可以参考[Link](https://jwt.io/libraries)

下面我使用了 `jjwt` 这个包

> <https://github.com/jwtk/jjwt>

```java
public interface JwtTokenProvider<T> {
    String generateToken(T details);

    T parsingToken(String token) throws InvalidTokenException;

    default void invalidToken(String token) {
        throw new UnsupportedOperationException();
    }
}

@Component("jjwt")
@Slf4j
@RequiredArgsConstructor
public class JJWT implements JwtTokenProvider<JwtUser> {
    private static final String USERNAME = "username";
    private static final String CACHE_KEY = "jjwt::invalid::token::";
    private final RedisTemplate<String, Object> redisTemplate;

    @Value("${app.security.jwt.secret-key:${random.uuid}}")
    private String secretKey;
    @Value("${application.security.jwt.expiration:10}")
    private long jwtExpiration;

    @Override
    public String generateToken(JwtUser user) {
        Date expireAt = new Date(System.currentTimeMillis() + TimeUnit.MILLISECONDS.convert(jwtExpiration, TimeUnit.MINUTES));

        return Jwts.builder()
                .claim(USERNAME, user.getUsername())
                .setExpiration(expireAt)
                .signWith(getSignInKey()).compact();
    }

    @Override
    public JwtUser parsingToken(String token) throws InvalidTokenException {
        if (isInvalidToken(token)) {
            throw new InvalidTokenException("Invalid token: " + token);
        }

        String username = parseClaimsJws(token).getBody().get(USERNAME, String.class);

        User user = new User();
        user.setUsername(username);
        return user;
    }

    @Override
    public void invalidToken(String token) {
        try {
            Date expirationDate = parseClaimsJws(token).getBody().getExpiration();
            long a = expirationDate.getTime() - System.currentTimeMillis();
            redisTemplate.opsForValue().set(getCacheKey(token), 1,
                    Duration.ofMillis(a));
        } catch (InvalidTokenException e) {
            log.warn("invalid token", e);
        }
    }

    private Jws<Claims> parseClaimsJws(String token) throws InvalidTokenException {
        try {
            return Jwts.parserBuilder()
                    .setSigningKey(getSignInKey()).build()
                    .parseClaimsJws(token);

        } catch (ExpiredJwtException e) {
            throw new TokenExpiredException("the token has expired", e);
        } catch (Exception e) {
            throw new InvalidTokenException("invalid token: " + token, e);
        }
    }

    private boolean isInvalidToken(String token) {
        return Boolean.TRUE.equals(redisTemplate.hasKey(getCacheKey(token)));
    }

    private String getCacheKey(String token) {
        return CACHE_KEY + token;
    }

    private Key getSignInKey() {
        return Keys.hmacShaKeyFor(secretKey.getBytes());
    }
}
```

## Handler

下面列举一些可以自定义的 Handler，这些 Handler 会在某些情况下执行

### AuthenticationSuccessHandler

这个是认证成功后的处理器，也就是在登录成功后执行。通过实现这个接口可以在登录成功后替我们做一些事情

下面的例子中在认证成功后向客户端响应 Token

```java
@Component
@Slf4j
@RequiredArgsConstructor
public class CustomizedAuthenticationSuccessHandler implements AuthenticationSuccessHandler {
    private final ObjectMapper objectMapper;
    private final JwtTokenProvider<JwtUser> tokenProvider;

    @Override
    public void onAuthenticationSuccess(HttpServletRequest request, HttpServletResponse response, Authentication authentication)
            throws IOException {
        JwtUser user = (JwtUser) authentication.getPrincipal();
        String token = tokenProvider.generateToken(user);

        response.setStatus(HttpStatus.OK.value());
        response.setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE);
        objectMapper.writeValue(response.getWriter(), Collections.singletonMap("token", token));
    }
}
```

### AuthenticationFailureHandler

登录认证失败处理器，就是账号密码不匹配时会执行

下面这个例子在认证失败后会向客户端响应错误信息

```java
@Component
@Slf4j
@RequiredArgsConstructor
public class CustomizedAuthenticationFailureHandler implements AuthenticationFailureHandler {
    private final ObjectMapper objectMapper;

    @Override
    public void onAuthenticationFailure(HttpServletRequest request, HttpServletResponse response,
                                        AuthenticationException exception) throws IOException {
        response.setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE);
        response.setStatus(HttpStatus.UNAUTHORIZED.value());
        objectMapper.writeValue(response.getWriter(), Collections.singletonMap("message", "incorrect username or password"));
    }
}

```

### AccessDeniedHandler

权限不足时的处理器，需要在用户已经通过认证了，但是访问的资源权限不够才会执行

```java
@Component
@Slf4j
@RequiredArgsConstructor
public class CustomizedAccessDeniedHandler implements AccessDeniedHandler {
    private final ObjectMapper objectMapper;

    @Override
    public void handle(HttpServletRequest request, HttpServletResponse response, AccessDeniedException accessDeniedException) throws IOException, ServletException {
        response.setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE);
        response.setStatus(HttpStatus.FORBIDDEN.value());
        objectMapper.writeValue(response.getWriter(), Collections.singletonMap("message", HttpStatus.FORBIDDEN.getReasonPhrase()));
    }
}
```

### AuthenticationEntryPoint

这个处理器是在用户未登录时直接访问受保护的资源时会执行

```java
@Component
@Slf4j
@RequiredArgsConstructor
public class CustomizedAuthenticationEntryPoint implements AuthenticationEntryPoint {
    private final ObjectMapper objectMapper;

    @Override
    public void commence(HttpServletRequest request, HttpServletResponse response, AuthenticationException authException) throws IOException, ServletException {
        response.setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE);
        response.setStatus(HttpStatus.UNAUTHORIZED.value());
        objectMapper.writeValue(response.getWriter(), Collections.singletonMap("message", "please login"));
    }
}
```

## UserDetailsService

这个接口时用来获取用户信息的，当将其实例注入容器时 Spring Security 在进行认证时就会通过这个接口获取用户信息来进行验证

```java
public interface UserDetailsService {
    UserDetails loadUserByUsername(String username) throws UsernameNotFoundException;
}
```

自定义实现

```java
@RequiredArgsConstructor
@Service
public class UserService implements UserDetailsService {
    private final UserRepo userRepo;

    @Override
    public JwtUser loadUserByUsername(String username) throws UsernameNotFoundException {
        return userRepo.findByUsername(username)
                .orElseThrow(() -> new UsernameNotFoundException("username not found"));
    }

    public Authentication getAuthentication(@Nonnull JwtUser user) {
        JwtUser dbUser = loadUserByUsername(user.getUsername());
        return UsernamePasswordAuthenticationToken.authenticated(dbUser, "", dbUser.getAuthorities());
    }
}
```

`InMemoryUserDetailsManager` 一个基于内存的 UserDetailsService 实现

## AuthenticationFilter

自定义一个 AuthenticationFilter 可以实现解析 Token，通过 **AuthenticationSuccessHandler** 我们已经可以将 Token 发送给客户端了。客户端使用 Token 访问资源时就需要通过这个 Filter 去解析 Token。

下面这个 Filter 实现了，从 `Authorization` 头中获取 Token，然后解析 Token，然后将解析出来的用户放入 `Security Context` 中。

```java
@Component
@Slf4j
@RequiredArgsConstructor
public class JwtTokenAuthenticationFilter extends OncePerRequestFilter {
    private final JwtTokenProvider<JwtUser> tokenProvider;
    private final UserService userService;


    protected String getToken(HttpServletRequest request) {
        String token = request.getHeader(HttpHeaders.AUTHORIZATION);
        if (Objects.isNull(token)) {
            return null;
        }
        return token.replaceFirst("Bearer ", "");
    }

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {

        String token = getToken(request);
        if (StringUtils.isNotEmpty(token)) {
            try {
                JwtUser jwtUser = tokenProvider.parsingToken(token);
                SecurityContextHolder.getContext().setAuthentication(userService.getAuthentication(jwtUser));
            } catch (TokenExpiredException e) {
                log.info("token expired: {}", token);
            } catch (InvalidTokenException e) {
                log.error("invalid token: {}", token, e);
            }
        }

        filterChain.doFilter(request, response);
    }
}

```

## 最后的总配置类

要想上面的组件生效还需要进行配置

```java
@Configuration
@EnableWebSecurity
@RequiredArgsConstructor
public class SecurityConfig {
    private final CustomizedAuthenticationSuccessHandler authenticationSuccessHandler;
    private final CustomizedAuthenticationFailureHandler authenticationFailureHandler;
    private final CustomizedAccessDeniedHandler accessDeniedHandler;
    private final CustomizedAuthenticationEntryPoint authenticationEntryPoint;
    private final JwtTokenAuthenticationFilter jwtTokenAuthenticationFilter;

    @Bean
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {
        http.csrf().disable()
                // 禁用 Session
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS);
        http
                .formLogin()
                // 配置表单登录接口
                .loginProcessingUrl("/api/login")
                // 认证成功处理器
                .successHandler(authenticationSuccessHandler)
                // 认证失败处理器
                .failureHandler(authenticationFailureHandler);

        http
                .authorizeHttpRequests()
                // 放行下面这个资源
                .requestMatchers("/api/resource1").permitAll()
                // 下面这个资源需要有 admin role 才可以访问
                // hasRole 会自动强制加上 ROLE_ 前缀，所以 user 中的 role 也必须要有 ROLE_前缀
                .requestMatchers("/api/resource3").hasRole("ADMIN")
                // 其他资源都需要认证
                .anyRequest().authenticated();

        http
                .exceptionHandling()
                // 权限不足处理器
                .accessDeniedHandler(accessDeniedHandler)
                // 未认证访问受保护资源处理器
                .authenticationEntryPoint(authenticationEntryPoint);

        http
                // 将 Token 解析的 Filter 加入 Filter Chain 中
                .addFilterBefore(jwtTokenAuthenticationFilter, UsernamePasswordAuthenticationFilter.class);

        http
                // 配置跨域
                .cors().configurationSource(request -> {
                    CorsConfiguration configuration = new CorsConfiguration();
                    configuration.setAllowedOrigins(Collections.singletonList(CorsConfiguration.ALL));
                    configuration.setAllowedMethods(Collections.singletonList(CorsConfiguration.ALL));
                    configuration.setAllowedHeaders(Collections.singletonList(CorsConfiguration.ALL));
                    return configuration;
                });

        return http.build();
    }

    /**
     * 指定密码加密器
     */
    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }
}
```

以上的某些配置并不是必须的
