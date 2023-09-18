class ReadFile {
    static void main(String[] args) {
        def file = new File("HelloWorld.txt")

        file.eachLine {
            line -> println "line:$line"
        }

        println "String: ${file.text}"

        println "The file ${file.absolutePath} has ${file.length()} bytes"

        println "File ? ${file.isFile()}"
        println "Directory ? ${file.isDirectory()}"

        println file.text[4]
        println file.text[2, 5]
        println file.text[2..5]
        println file.text[2..<5]
        println file.text['a'..'c']
        println file.text[6..2]
    }
}
