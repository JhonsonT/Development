class HelloWorld {
    static void main(String[] args) {
        String hello = "hello"
        def world = "world"

        def range = 5..10

        println(hello)
        println(world)
        println(hello + " " + world)
        println(range)
        println(range.get(2))

        test(range, range.get(1))
        test(range.get(3))
    }

    static def test(p1, p2 = 0) {
        println(p1)
        println(p2)
    }
}
