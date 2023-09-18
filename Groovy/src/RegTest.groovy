class RegTest {
    static void main(String[] args) {
        def regex = 'Groovy' =~ 'Groovy'
        println regex.matches()

        regex = 'Groovy' =~ 'oo'
        println regex.matches()

        regex = 'Groovy' =~ '^G'
        println regex.matches()

        regex = 'Groovy' =~ 'G$'
        println regex.matches()

        regex = 'Groovy' =~ 'Gro*vy'
        println regex.matches()

        regex = 'Groovy' =~ 'Gro{2}vy'
        println regex.matches()
    }
}
