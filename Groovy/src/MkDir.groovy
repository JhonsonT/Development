class MkDir {
    static void main(String[] args) {
        println("user.home：" + System.getProperty("user.home"))

        String path = System.getProperty("user.home") + "/Downloads/test";
        println path
        def file = new File(path)
        file.mkdir()

        file.delete()
    }
}
