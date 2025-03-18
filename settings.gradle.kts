rootProject.name = "proto-specs"

// Automatically include all subprojects (Java and Kotlin)
rootDir.listFiles()
    ?.filter { it.isDirectory && File(it, "build.gradle.kts").exists() }
    ?.forEach { include(it.name) }
findProject("kotlin")?.name = "proto-kotlin"
findProject("java")?.name = "proto-java"