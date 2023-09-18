class ClosureTest {
    def static closureTest(clo) {
        clo.call("sx，呵呵")
    }

    static void main(String[] args) {
        def s = '大'

        def closure = { param -> println "${s}${param}，你好" }
        closure.call("sx")

        closure = { println "${s}${it}，你好" }
        closure.call("sx")

        closureTest(closure)
    }
}
