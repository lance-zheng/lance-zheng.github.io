<!-- customize-category:Java -->

# MapStruct

- [MapStruct](#mapstruct)
  - [Mapping Customization](#mapping-customization)
    - [Decorator Pattern](#decorator-pattern)

<https://github.com/mapstruct/mapstruct>

## Mapping Customization

有时候某些需求可能么法通过 MapStruct 做到，我们需要在转换前后执行一些自定义逻辑。

### Decorator Pattern

可以使用 MapStruct 中提供的 `@DecoratedWith` 注解。

```java
public abstract class UserDtoConverterDecorator implements UserDtoConverter {
    @Autowired
    @Qualifier("delegate")
    private  UserDtoConverter delegate;

    @Override
    public User toEntity(UserDto dto) {
        return delegate.toEntity(dto);
    }

    @Override
    public void updateEntity(UserDto dto, User entity) {
        delegate.updateEntity(dto, entity);
    }

    @Override
    public UserDto toDto(User entity) {
        return delegate.toDto(entity);
    }
}


@Mapper(componentModel = "spring")
@DecoratedWith(UserDtoConverterDecorator.class)
interface UserDtoConverter extends DtoConverter<User, UserDto> {
    @Override
    @Mapping(target = "password", ignore = true)
    UserDto toDto(User entity);
}
```

这里我们使用了 `@Mapper(componentModel = "spring")` Spring 容器，在生成的类中会自动 `@Component` 注解。

生成的文件是这样的。

```java
@Component
@Primary
class UserDtoConverterImpl extends UserDtoConverterDecorator {
}

@Component
@Qualifier("delegate")
class UserDtoConverterImpl_ implements UserDtoConverter {
    //...
}
```

下面这种是不使用 `@DecoratedWith` 注解的写法。

```java
@Component
public class UserDtoConverterDecorator implements UserDtoConverter {
    private final UserDtoConverter delegate = Mappers.getMapper(UserDtoConverter.class);

    @Override
    public User toEntity(UserDto dto) {
        return delegate.toEntity(dto);
    }

    @Override
    public void updateEntity(UserDto dto, User entity) {
        delegate.updateEntity(dto, entity);
    }

    @Override
    public UserDto toDto(User entity) {
        return delegate.toDto(entity);
    }
}


@Mapper
interface UserDtoConverter extends DtoConverter<User, UserDto> {
    @Override
    @Mapping(target = "password", ignore = true)
    UserDto toDto(User entity);
}
```
