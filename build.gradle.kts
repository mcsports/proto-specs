plugins {
    id("com.google.protobuf") version "0.9.6" apply false
}

allprojects {

    group = "club.mcsports.generated"
    version = findProperty("version") ?: "1.0.0"

    repositories {
        mavenCentral()
    }
}