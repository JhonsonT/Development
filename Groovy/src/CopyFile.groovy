class CopyFile {
    static void main(String[] args) {
        def file = new File("HelloWorld.txt")
        def file2 = new File("Hello.txt")
        file2 << file.text
    }
}
