plugins {
    alias(libs.plugins.protobuf) apply false
}

allprojects {

    group = "club.mcsports.generated"
    version = findProperty("version") ?: "1.0.0"

    repositories {
        mavenCentral()
    }
}