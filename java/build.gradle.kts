plugins {
    id("java")
    id("com.google.protobuf") version "0.9.4"
    id("maven-publish")
}

java {
    sourceCompatibility = JavaVersion.VERSION_21
}

dependencies {
    implementation("com.google.protobuf:protobuf-java:4.31.0")
    implementation("io.grpc:grpc-protobuf:1.71.0")
    implementation("io.grpc:grpc-stub:1.71.0")
    implementation("javax.annotation:javax.annotation-api:1.3.2")
}

protobuf {
    protoc { artifact = "com.google.protobuf:protoc:4.30.1" }
    plugins {
        register("grpc") { artifact = "io.grpc:protoc-gen-grpc-java:1.71.0" }
    }
    generateProtoTasks {
        all().configureEach {
            plugins { register("grpc") }
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
