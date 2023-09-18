import groovy.xml.MarkupBuilder
import groovy.xml.XmlParser

class XmlTest {
    static void main(String[] args) {
        def builder = new MarkupBuilder()

        builder.xml(st: '呵呵') {
            yy(ss: 'as') {
                sa('22')
            }
        }

        def mp = [1: ["11", "22"], 2: ["223", "333"]]

        builder.xml(ss: '啦啦啦') {
            mp.each {
                demo ->
                    {
                        builder.yy(21: demo.value[0]) {
                            simple(demo.value[1])
                        }
                    }
            }
        }

        def parser = new XmlParser()
        def parse = parser.parse("XmlDemo")
        parse.yy.each{
            println it.sa
            println it.simple[0]
        }
    }
}
