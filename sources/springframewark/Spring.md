<!-- customize-category:Spring Framework -->

# Bean 的生命周期

```markdown
1. Construct # 实例化对象
2. @Autowired # 依赖注入
3. BeanNameAware # 执行 Aware 相关函数
   BeanFactoryAware
   ApplicationContextAware
4. BeanPostProcessor#postProcessBeforeInitialization() # 执行 BeanPostProcessor 的前置处理函数
   @PostConstruct # 这个也是通过 BeanPostProcess 前置处理实现的
5. InitializingBean#afterPropertiesSet #
6. initMethod # 通过 @Bean(initMethod = "initMethod") 可以指定
7. BeanPostProcessor#postProcessAfterInitialization()
8. @PreDestroy
9. DisposableBean#destroy
10. destroyMethod # 通过 @Bean(destroyMethod = "destroyMethod") 可以指定
```

```java
// @Component
public class MyBean implements InitializingBean, BeanNameAware, ApplicationContextAware, BeanFactoryAware, DisposableBean {
    MyBeanPostProcess myBeanPostProcess;

    @Autowired
    public void setMyBeanPostProcess(MyBeanPostProcess myBeanPostProcess) {
        System.out.println("@Autowired");
        this.myBeanPostProcess = myBeanPostProcess;
    }

    public MyBean() {
        System.out.println("Construct");
    }

    // InitDestroyAnnotationBeanPostProcessor
    @PostConstruct
    public void postConstruct() {
        System.out.println("@PostConstruct");
    }

    // InitDestroyAnnotationBeanPostProcessor
    @PreDestroy
    public void preDestroy() {
        System.out.println("@PreDestroy");
    }

    @Override
    public void afterPropertiesSet() throws Exception {
        System.out.println("InitializingBean#afterPropertiesSet");
    }

    @Override
    public void destroy() throws Exception {
        System.out.println("DisposableBean#destroy");
    }

    @Override
    public void setBeanName(String name) {
        System.out.println("BeanNameAware");
    }

    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        System.out.println("ApplicationContextAware");
    }

    @Override
    public void setBeanFactory(BeanFactory beanFactory) throws BeansException {
        System.out.println("BeanFactoryAware");
    }

    public void initMethod() {
        System.out.println("initMethod");
    }

    public void destroyMethod() {
        System.out.println("destroyMethod");
    }


}

@Component
class MyBeanPostProcess implements BeanPostProcessor {
    @Override
    public Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException {
        if (bean instanceof MyBean) {
            System.out.println("BeanPostProcessor#postProcessBeforeInitialization()");
        }
        return BeanPostProcessor.super.postProcessBeforeInitialization(bean, beanName);
    }

    @Override
    public Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException {
        if (bean instanceof MyBean) {
            System.out.println("BeanPostProcessor#postProcessAfterInitialization()");
        }
        return BeanPostProcessor.super.postProcessAfterInitialization(bean, beanName);
    }

    @Bean(initMethod = "initMethod", destroyMethod = "destroyMethod")
    MyBean foo() {
        return new MyBean();
    }
}
```
