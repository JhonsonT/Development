class WriteFile {
    static void main(String[] args) {
        def file1 = new File("HelloWorld.txt")
                .withWriter('utf-8')
                        {
                            w -> w.writeLine('Hello World')
                        }
    }
}
