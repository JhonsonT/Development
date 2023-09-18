import groovy.transform.AnnotationCollector

@Tree(str = "11", str2 = "22")
class AnnotationTest {
    static void main(String[] args) {

    }
}

@interface One {
    String str() default "Hello World"
}

@interface Two {
    String str2() default "Hello World"
}

@One
@Two
@AnnotationCollector
class Tree {

}
