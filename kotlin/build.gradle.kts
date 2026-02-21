plugins {
    id("org.jetbrains.kotlin.jvm") version "2.3.10"
    id("com.google.protobuf") version "0.9.6"
    id("maven-publish")
}

dependencies {
    api(project(":proto-java"))
    implementation("com.google.protobuf:protobuf-kotlin:4.33.5")
    implementation("io.grpc:grpc-protobuf:1.79.0")
    implementation("io.grpc:grpc-stub:1.79.0")
    implementation("io.grpc:grpc-kotlin-stub:1.5.0")
}

protobuf {
    protoc { artifact = "com.google.protobuf:protoc:4.33.4" }
    plugins {
        register("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:1.79.0"
        }
        register("grpckt") { artifact = "io.grpc:protoc-gen-grpc-kotlin:1.79.0" }
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
        val reposiliteUrl = findProperty("reposiliteUrl") as String?
        if(reposiliteUrl != null) {
            maven {
                name = "Reposilite"
                url = uri(reposiliteUrl)
                credentials {
                    username = findProperty("reposiliteUsername") as String?
                    password = findProperty("reposilitePassword") as String?
                }
            }
        }
    }

    publications {
        create<MavenPublication>("maven") {
            from(components["java"])
        }
    }
}