plugins {
    id("org.jetbrains.kotlin.jvm") version "2.0.20"
    id("com.google.protobuf")
    id("maven-publish")
}

dependencies {
    implementation(project(":proto-java"))
    implementation("com.google.protobuf:protobuf-kotlin:4.30.1")
    implementation("io.grpc:grpc-protobuf:1.71.0")
    implementation("io.grpc:grpc-stub:1.71.0")
    implementation("io.grpc:grpc-kotlin-stub:1.4.1")
}

protobuf {
    protoc { artifact = "com.google.protobuf:protoc:4.30.1" }
    plugins {
        register("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:1.71.0"
        }
        register("grpckt") { artifact = "io.grpc:protoc-gen-grpc-kotlin:1.71.0" }
    }
    generateProtoTasks {
        all().configureEach {
            plugins {
                register("grpc")
                register("grpckt")
            }
        }
    }
}

publishing {
    repositories {
        maven {
            name = "Reposilite"
            url = uri(findProperty("reposiliteUrl") ?: "")
            credentials {
                username = findProperty("reposiliteUsername") as String?
                password = findProperty("reposilitePassword") as String?
            }
        }
    }

    publications {
        create<MavenPublication>("maven") {
            from(components["java"])
        }
    }
}