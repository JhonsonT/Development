class ReadDir {
    static void main(String[] args) {
        def rootFiles = new File(System.getProperty("user.home")).listFiles()
        rootFiles.each {
            file -> println file.absolutePath
        }

        println "*******************************"

        new File(System.getProperty("user.home") + "/Documents")
                .eachFile { println it.absolutePath }

        println "*******************************"

        new File(System.getProperty("user.home") + "/Documents/OCC")
                .eachFileRecurse { println it.absolutePath }
    }
}
