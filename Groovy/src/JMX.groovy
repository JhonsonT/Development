import java.lang.management.ManagementFactory

def os = ManagementFactory.operatingSystemMXBean
println os.arch
println os.name
println os.version
println os.availableProcessors