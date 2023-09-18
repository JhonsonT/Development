class TraitTest implements Hello {

    static void main(String[] args) {
        TraitTest test = new TraitTest();
        test.sayHello()
    }
}

interface Say {
    sayHello()
}

trait Hello implements Say {
    def sayHello() {
        println 'Hello'
    }
}
