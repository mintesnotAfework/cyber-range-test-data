package models

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"minta/docker"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func test() {
	PullImage()
	adminTest()
	instructorTest()
	studentTest()
	roomInstructorTest()
	roomStudentTest()
	courseTest()
	machineTest()
	questionTest()
	questionStudentTest()
	notificationTest()
}

func adminTest() {
	value, _ := bcrypt.GenerateFromPassword([]byte("SecureP@ssw0rd"), 12)
	user := User{
		UserName:      "Admin",
		FirstName:     "admin",
		LastName:      "admin",
		MiddleName:    "admin",
		Email:         "admin.admin@admin.com",
		Phone:         "+25199090909",
		Password:      string(value),
		Locked:        false,
		EmailVerified: true,
		PhoneVerified: true,
	}

	user.CreateUser()

	admin := Admin{
		UserId: 1,
	}
	_, err := admin.CreateAdmin()
	if err != nil {
		log.Println(err.Error())
	}
}

func instructorTest() {
	value, _ := bcrypt.GenerateFromPassword([]byte("SecureP@ssw0rd"), 12)
	users := []User{
		{
			UserName:      "bob_jones4",
			FirstName:     "Bob",
			LastName:      "Jones",
			MiddleName:    "D",
			Email:         "bob.jones@example.com",
			Phone:         "+251911223377",
			Password:      string(value),
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		},
		{
			UserName:      "charlie_brown5",
			FirstName:     "Charlie",
			LastName:      "Brown",
			MiddleName:    "E",
			Email:         "charlie.brown@example.com",
			Phone:         "+251911223388",
			Password:      string(value),
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		},
		{
			UserName:      "david_clark6",
			FirstName:     "David",
			LastName:      "Clark",
			MiddleName:    "F",
			Email:         "david.clark@example.com",
			Phone:         "+251911223399",
			Password:      string(value),
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		},
	}
	for _, user := range users {
		cu, err := user.CreateUser()
		if err != nil {
			return
		}

		instructor := &Instructor{
			AccountVerified: true,
			UserId:          cu.ID,
		}

		instructor.CreateInstructor()
	}

	for i := 0; i < 50; i++ {
		phone := 100000000
		user := User{
			UserName:      generateRandomString(10),
			FirstName:     generateRandomString(7),
			LastName:      generateRandomString(7),
			MiddleName:    generateRandomString(5),
			Email:         fmt.Sprintf("instructor%d@example.com", i+1),
			Phone:         fmt.Sprintf("+%d", phone+i), // Last two digits increment
			Password:      string(value),               // Assuming value is defined elsewhere
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		}
		cu, err := user.CreateUser()
		if err != nil {
			log.Println(err.Error())
		}
		instructor := Instructor{
			AccountVerified: true,
			UserId:          cu.ID, // Assuming the first 7 users are already created
		}
		_, err = instructor.CreateInstructor()
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func studentTest() {
	value, _ := bcrypt.GenerateFromPassword([]byte("SecureP@ssw0rd"), 12)
	users := []User{
		{
			UserName:      "john_doe1",
			FirstName:     "John",
			LastName:      "Doe",
			MiddleName:    "A",
			Email:         "john.doe@example.com",
			Phone:         "+251911223344",
			Password:      string(value),
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		},
		{
			UserName:      "jane_doe2",
			FirstName:     "Jane",
			LastName:      "Doe",
			MiddleName:    "B",
			Email:         "jane.doe@example.com",
			Phone:         "+251911223355",
			Password:      string(value),
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		},
		{
			UserName:      "alice_smith3",
			FirstName:     "Alice",
			LastName:      "Smith",
			MiddleName:    "C",
			Email:         "alice.smith@example.com",
			Phone:         "+251911223366",
			Password:      string(value),
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		},
	}

	for _, user := range users {
		cu, err := user.CreateUser()
		if err != nil {
			return
		}
		student := &Student{
			IsPremiem: false,
			UserId:    cu.ID,
		}

		student.CreateStudent()
	}

	for i := 0; i < 50; i++ {
		phone := 200000000
		user := User{
			UserName:      generateRandomString(10),
			FirstName:     generateRandomString(7),
			LastName:      generateRandomString(7),
			MiddleName:    generateRandomString(5),
			Email:         fmt.Sprintf("student%d@example.com", i+1),
			Phone:         fmt.Sprintf("+%d", phone+i), // Last two digits increment
			Password:      string(value),               // Assuming value is defined elsewhere
			Locked:        false,
			EmailVerified: true,
			PhoneVerified: true,
		}
		cu, err := user.CreateUser()
		if err != nil {
			log.Println(err.Error())
		}

		student := Student{
			IsPremiem: false,
			UserId:    cu.ID, // Assuming the first 7 users are already created
		}
		_, err = student.CreateStudent()
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func roomInstructorTest() {
	rooms := []Room{
		{
			Title:        "Cybersecurity Fundamentals",
			Description:  "Introduction to cybersecurity concepts and principles",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       false,
		},
		{
			Title:        "Network Security",
			Description:  "Securing network infrastructure and protocols",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       true,
		},
		{
			Title:        "Web Application Security",
			Description:  "Protecting web applications from vulnerabilities",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       false,
		},
		{
			Title:        "Cryptography",
			Description:  "Encryption algorithms and secure communication",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       true,
		},
		{
			Title:        "Penetration Testing",
			Description:  "Ethical hacking and vulnerability assessment",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       false,
		},
		{
			Title:        "Digital Forensics",
			Description:  "Investigating cyber crimes and analyzing digital evidence",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       true,
		},
		{
			Title:        "Secure Coding Practices",
			Description:  "Writing code resistant to vulnerabilities",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       true,
		},
		{
			Title:        "Cloud Security",
			Description:  "Securing cloud infrastructure and services",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       false,
		},
		{
			Title:        "Malware Analysis",
			Description:  "Understanding and reverse-engineering malicious software",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       true,
		},
		{
			Title:        "Incident Response",
			Description:  "Handling and mitigating security breaches",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       false,
		},
		{
			Title:        "Advanced Threat Intelligence",
			Description:  "Analyzing and anticipating cyber threats",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       false,
		},
		{
			Title:        "IoT Security",
			Description:  "Securing Internet of Things devices and networks",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       true,
		},
		{
			Title:        "Security Operations Center",
			Description:  "Monitoring and defending organizational networks",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       false,
		},
		{
			Title:        "Social Engineering Defense",
			Description:  "Protecting against human-factor exploits",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       true,
		},
		{
			Title:        "Blockchain Security",
			Description:  "Security aspects of blockchain technology",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       false,
		},
		{
			Title:        "Ethical Hacking",
			Description:  "Authorized penetration testing methodologies",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       false,
		},
		{
			Title:        "Advanced Persistent Threats",
			Description:  "Understanding and defending against APTs",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       true,
		},
		{
			Title:        "Mobile Security",
			Description:  "Securing mobile devices and applications",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       false,
		},
		{
			Title:        "Red Team Operations",
			Description:  "Simulating adversary attacks for defense",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       true,
		},
		{
			Title:        "Blue Team Defense",
			Description:  "Implementing defensive security measures",
			InstructorId: 1,
			RoomVerified: false,
			Locked:       false,
		},
		{
			Title:        "Security Compliance",
			Description:  "Understanding regulatory security requirements",
			InstructorId: 1,
			RoomVerified: true,
			Locked:       false,
		},
	}
	for index, room := range rooms {
		room.CreatedAt = time.Now().AddDate(0, -1*index, 0)
		_, err := room.CreateRoom()
		if err != nil {
			log.Println(err.Error())
		}
	}
	for i := 0; i < 50; i++ {
		room := Room{
			Title:        "Ethical Hacking And Penteration Testing Part " + strconv.Itoa(i),
			Description:  generateRandomString(30),
			InstructorId: 1,
			RoomVerified: true,
			Locked:       false,
			CreatedAt:    time.Now().AddDate(-10, 0, 0),
		}
		_, err := room.CreateRoom()
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func roomStudentTest() {
	roomStudents := []RoomStudent{
		{
			MemberId: 1,
			RoomId:   1,
		},
		{
			MemberId: 2,
			RoomId:   2,
		},
		{
			MemberId: 2,
			RoomId:   1,
		},
		{
			MemberId: 3,
			RoomId:   1,
		},
		{
			MemberId: 3,
			RoomId:   3,
		},
		{
			MemberId: 1,
			RoomId:   4,
		},
		{
			MemberId: 2,
			RoomId:   5,
		},
		{
			MemberId: 3,
			RoomId:   6,
		},
	}
	for _, roomStudent := range roomStudents {
		_, err := roomStudent.CreateRoomStudent()
		if err != nil {
			log.Println(err.Error())
		}
	}
	for i := 1; i > 100; i = i + 3 {
		roomStudent := RoomStudent{
			MemberId: 1,
			RoomId:   uint(i),
		}
		_, err := roomStudent.CreateRoomStudent()
		if err != nil {

		}
	}
	for i := 2; i > 100; i = i + 3 {
		roomStudent := RoomStudent{
			MemberId: 2,
			RoomId:   uint(i),
		}
		_, err := roomStudent.CreateRoomStudent()
		if err != nil {

		}
	}
	for i := 3; i > 100; i = i + 3 {
		roomStudent := RoomStudent{
			MemberId: 3,
			RoomId:   uint(i),
		}
		_, err := roomStudent.CreateRoomStudent()
		if err != nil {

		}
	}
}

func courseTest() {
	courseMachines1 := []CourseMachine{
		{
			Title:             "Weak Password Policies",
			Description:       "Exploring the risks of simple passwords and poor credential management",
			Point:             150,
			DifficultyLevelId: 2,
		},
		{
			Title:             "Phishing Attack Simulation",
			Description:       "How attackers trick users via deceptive emails and links",
			Point:             200,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Unpatched Software Vulnerabilities",
			Description:       "The dangers of failing to update systems and applications",
			Point:             180,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Man-in-the-Middle Attacks",
			Description:       "How attackers intercept and manipulate network communications",
			Point:             250,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Insecure Wi-Fi Networks",
			Description:       "Risks of using public or poorly secured wireless networks",
			Point:             150,
			DifficultyLevelId: 2,
		},
		{
			Title:             "Social Engineering Exploits",
			Description:       "Psychological manipulation to gain unauthorized access",
			Point:             200,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Ransomware Infection Scenario",
			Description:       "How ransomware encrypts files and demands payment",
			Point:             99,
			DifficultyLevelId: 5,
		},
		{
			Title:             "SQL Injection Attacks",
			Description:       "Exploiting database vulnerabilities through malicious queries",
			Point:             250,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Insider Threats And Data Leaks",
			Description:       "Risks posed by employees mishandling sensitive data",
			Point:             200,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Zero-Day Exploit Awareness",
			Description:       "Understanding unknown vulnerabilities before patches exist",
			Point:             90,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Firewall Misconfigurations",
			Description:       "Common firewall rule errors that expose networks to attacks",
			Point:             180,
			DifficultyLevelId: 3,
		},
		{
			Title:             "DDoS Attack Mitigation",
			Description:       "Defending against Distributed Denial of Service attacks",
			Point:             250,
			DifficultyLevelId: 4,
		},
		{
			Title:             "VPN Security Flaws",
			Description:       "Exploiting weaknesses in Virtual Private Networks",
			Point:             220,
			DifficultyLevelId: 4,
		},
		{
			Title:             "ARP Spoofing Attacks",
			Description:       "How attackers manipulate Address Resolution Protocol",
			Point:             200,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Wi-Fi Eavesdropping",
			Description:       "Intercepting unencrypted wireless communications",
			Point:             170,
			DifficultyLevelId: 2,
		},
		{
			Title:             "Network Intrusion Detection",
			Description:       "Identifying malicious activity using NIDS/NIPS",
			Point:             230,
			DifficultyLevelId: 4,
		},
		{
			Title:             "DNS Hijacking",
			Description:       "Redirecting traffic through compromised DNS servers",
			Point:             210,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Zero Trust Architecture",
			Description:       "Implementing 'never trust, always verify' principles",
			Point:             240,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Rogue Access Points",
			Description:       "Detecting and neutralizing unauthorized wireless devices",
			Point:             190,
			DifficultyLevelId: 3,
		},
		{
			Title:             "BGP Hijacking",
			Description:       "Exploiting Border Gateway Protocol vulnerabilities",
			Point:             199,
			DifficultyLevelId: 5,
		},
		{
			Title:             "SQL Injection",
			Description:       "Exploiting database queries through input fields",
			Point:             250,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Cross-Site Scripting (XSS)",
			Description:       "Injecting malicious scripts into web pages",
			Point:             220,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Cross-Site Request Forgery (CSRF)",
			Description:       "Forcing users to execute unwanted actions",
			Point:             230,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Broken Authentication",
			Description:       "Exploiting weak session management",
			Point:             200,
			DifficultyLevelId: 3,
		},
		{
			Title:             "API Security Misconfigurations",
			Description:       "Exploiting improperly secured REST/SOAP endpoints",
			Point:             240,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Weak Encryption Algorithms",
			Description:       "Breaking deprecated ciphers (DES, RC4)",
			Point:             220,
			DifficultyLevelId: 3,
		},
		{
			Title:             "RSA Key Factorization",
			Description:       "Cracking short RSA keys (<= 1024 bits)",
			Point:             210,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Hash Collision Attacks",
			Description:       "Exploiting MD5/SHA-1 weaknesses",
			Point:             120,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Padding Oracle Attacks",
			Description:       "Decrypting data via CBC padding errors",
			Point:             130,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Quantum Cryptography Basics",
			Description:       "Understanding post-quantum encryption",
			Point:             140,
			DifficultyLevelId: 5,
		},
		{
			Title:             "OSINT Reconnaissance",
			Description:       "Gathering target info from public sources",
			Point:             180,
			DifficultyLevelId: 2,
		},
		{
			Title:             "Privilege Escalation",
			Description:       "Gaining admin rights from low-level access",
			Point:             250,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Post-Exploitation Techniques",
			Description:       "Maintaining access after initial compromise",
			Point:             222,
			DifficultyLevelId: 5,
		},
		{
			Title:             "AD Domain Compromise",
			Description:       "Attacking Active Directory environments",
			Point:             98,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Red Team Operations",
			Description:       "Simulating advanced adversaries",
			Point:             78,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Memory Dump Analysis",
			Description:       "Extracting artifacts from RAM captures",
			Point:             230,
			DifficultyLevelId: 4,
		},
		{
			Title:             "File Carving Techniques",
			Description:       "Recovering deleted/hidden files",
			Point:             210,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Timeline Analysis",
			Description:       "Reconstructing attack sequences",
			Point:             240,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Anti-Forensics Detection",
			Description:       "Identifying evidence tampering",
			Point:             60,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Mobile Device Forensics",
			Description:       "Extracting data from smartphones",
			Point:             250,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Input Validation",
			Description:       "Preventing injection attacks",
			Point:             200,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Memory Safety",
			Description:       "Avoiding buffer overflows (C/C++)",
			Point:             240,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Secure API Design",
			Description:       "Implementing OAuth2/OpenID Connect",
			Point:             230,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Cryptographic Implementations",
			Description:       "Avoiding common crypto pitfalls",
			Point:             250,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Supply Chain Security",
			Description:       "Securing third-party dependencies",
			Point:             220,
			DifficultyLevelId: 3,
		},
		{
			Title:             "IAM Privilege Escalation",
			Description:       "Exploiting AWS/GCP/Azure permissions",
			Point:             26,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Container Breakouts",
			Description:       "Escaping Docker/Kubernetes isolation",
			Point:             20,
			DifficultyLevelId: 5,
		},
		{
			Title:             "Storage Bucket Misconfigurations",
			Description:       "Exploiting public S3/Blob Storage",
			Point:             10,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Serverless Function Abuse",
			Description:       "Exploiting AWS Lambda/Azure Functions",
			Point:             240,
			DifficultyLevelId: 4,
		},
		{
			Title:             "Cloud Logging Bypass",
			Description:       "Evading CloudTrail/Azure Monitor",
			Point:             250,
			DifficultyLevelId: 4,
		},
	}
	file, err := os.ReadFile("./models/courseContentTesting.md")
	if err != nil {
		return
	}

	content := string(file)

	for _, courseMachine := range courseMachines1 {
		cm, err := courseMachine.CreateCourseMachine()
		if err != nil {
			log.Println(err.Error())
		}
		course := Course{
			RoomId:          1, // Assuming room ID 1 for testing
			CourseMachineId: cm.ID,
			Content:         content,
		}
		course.CreateCourse()
	}
	for i := 0; i < 50; i++ {
		courseMachine := CourseMachine{
			Title:             "Cyber security and " + generateRandomString(5),
			Description:       generateRandomString(100),
			Point:             100,
			DifficultyLevelId: 2,
		}
		cm, err := courseMachine.CreateCourseMachine()
		if err != nil {
			log.Println(err.Error())
		}
		course := Course{
			RoomId:          2, // Assuming room ID 1 for testing
			CourseMachineId: cm.ID,
			Content:         content,
		}
		course.CreateCourse()
	}
}

func machineTest() {
	courseMachines := []CourseMachine{
		{
			Title:             "SQLi Fundamentals",
			Description:       "Learn and exploit basic SQL Injection vulnerabilities in a simulated environment.",
			Point:             150,
			DifficultyLevelId: 5,
		},
		{
			Title:             "XSS Playground",
			Description:       "Explore Cross-Site Scripting (XSS) attacks and how to prevent them.",
			Point:             120,
			DifficultyLevelId: 1,
		},
		{
			Title:             "Command Injection 101",
			Description:       "Understand command injection vulnerabilities through hands-on exploitation.",
			Point:             180,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Broken Auth Basics",
			Description:       "Learn to identify and exploit broken authentication mechanisms.",
			Point:             170,
			DifficultyLevelId: 2,
		},
		{
			Title:             "File Inclusion Lab",
			Description:       "Experiment with Local and Remote File Inclusion vulnerabilities.",
			Point:             190,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Sensitive Data Exposure",
			Description:       "Practice identifying sensitive data leaks in web applications.",
			Point:             140,
			DifficultyLevelId: 2,
		},
		{
			Title:             "Privilege Escalation Basics",
			Description:       "Simulate escalating privileges on a compromised machine.",
			Point:             200,
			DifficultyLevelId: 3,
		},
		{
			Title:             "Insecure Deserialization",
			Description:       "Explore the dangers of insecure deserialization and gain RCE.",
			Point:             220,
			DifficultyLevelId: 4,
		},
		{
			Title:             "JWT Hacking",
			Description:       "Learn how JSON Web Tokens (JWTs) can be manipulated for unauthorized access.",
			Point:             160,
			DifficultyLevelId: 2,
		},
		{
			Title:             "Web Shell Challenge",
			Description:       "Upload and use a web shell to compromise the system.",
			Point:             210,
			DifficultyLevelId: 4,
		},
	}

	images := []string{
		"mintesnotafework/sql_injection-image:v1.0",
		"mintesnotafework/xss_playground-image:v1.0",
		"mintesnotafework/command-injection-image:v1.0",
		"mintesnotafework/broken-auth-image:v1.0",
		"mintesnotafework/file-inclusion-image:v1.0",
		"mintesnotafework/sensitive_data_exposure-image:v1.0",
		"mintesnotafework/sql_injection-image:v1.0",
		"mintesnotafework/insecure-deserialization-image:v1.0",
		"mintesnotafework/sql_injection-image:v1.0",
		"mintesnotafework/sql_injection-image:v1.0",
	}

	for index, courseMachine := range courseMachines {
		cm, err := courseMachine.CreateCourseMachine()
		if err != nil {
			log.Println(err.Error())
		}

		machine := Machine{
			RoomId:                1,
			CourseMachineId:       cm.ID,
			ImageNameOrID:         images[index],
			OperatingSystemTypeId: 1,
		}

		_, err = machine.CreateMachine()
		if err != nil {
			log.Println(err.Error())
			continue
		}

	}
}

func questionTest() {
	questions := []Question{
		{
			CourseId:          1,
			Question:          "What is the basic unit of electric current?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 1,
			Point:             100,
		},
		{
			CourseId:          1,
			Question:          "What is the speed of light in vacuum?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 2,
			Point:             200,
		},
		{
			CourseId:          1,
			Question:          "Who is known as the father of modern physics?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 3,
			Point:             50,
		},
		{
			CourseId:          1,
			Question:          "What is the formula for calculating force?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 4,
			Point:             150,
		},
		{
			CourseId:          1,
			Question:          "What is the first law of thermodynamics?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 5,
			Point:             30,
		},
		{
			CourseId:          1,
			Question:          "What is the unit of frequency?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 1,
			Point:             99,
		},
		{
			CourseId:          1,
			Question:          "What is the principle of superposition?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 2,
			Point:             110,
		},
		{
			CourseId:          1,
			Question:          "What is the basic unit of electric current?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 1,
			Point:             100,
		},
		{
			CourseId:          1,
			Question:          "Who is known as the father of modern physics?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 3,
			Point:             50,
		},
		{
			CourseId:          1,
			Question:          "What is the speed of light in vacuum?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 2,
			Point:             200,
		},
		{
			CourseId:          1,
			Question:          "What is the formula for calculating force?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 4,
			Point:             150,
		},
		{
			CourseId:          1,
			Question:          "What is the first law of thermodynamics?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 5,
			Point:             30,
		},
		{
			CourseId:          1,
			Question:          "What is the unit of frequency?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 1,
			Point:             99,
		},
		{
			CourseId:          1,
			Question:          "What is the principle of superposition?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 2,
			Point:             110,
		},
		{
			CourseId:          1,
			Question:          "What is the basic unit of electric current?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 1,
			Point:             100,
		},
		{
			CourseId:          1,
			Question:          "What is the speed of 1234?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 2,
			Point:             200,
		},
		{
			CourseId:          1,
			Question:          "Who is known as the father of modern physics?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 3,
			Point:             50,
		},
		{
			CourseId:          1,
			Question:          "What is the formula for calculating force?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 4,
			Point:             150,
		},
		{
			CourseId:          1,
			Question:          "What is the first law of thermodynamics?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 5,
			Point:             30,
		},
		{
			CourseId:          1,
			Question:          "What is the unit of frequency?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 1,
			Point:             99,
		},
		{
			CourseId:          1,
			Question:          "What is the principle of superposition?",
			Answer:            "1234",
			Hint1:             "hello world",
			Hint2:             "hello world 2",
			Hint3:             "hello world 3",
			DifficultyLevelId: 2,
			Point:             110,
		},
	}

	for _, question := range questions {
		_, err := question.CreateQuestion()
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func questionStudentTest() {
	machineStudents := []QuestionStudent{
		{
			StudentId:  1,
			QuestionId: 1,
		},
		{
			StudentId:  1,
			QuestionId: 2,
		},
		{
			StudentId:  1,
			QuestionId: 3,
		},
		{
			StudentId:  1,
			QuestionId: 4,
		},
		{
			StudentId:  1,
			QuestionId: 5,
		},
		{
			StudentId:  2,
			QuestionId: 1,
		},
		{
			StudentId:  2,
			QuestionId: 2,
		},
		{
			StudentId:  2,
			QuestionId: 3,
		},
		{
			StudentId:  2,
			QuestionId: 4,
		},
		{
			StudentId:  2,
			QuestionId: 5,
		},
		{
			StudentId:  2,
			QuestionId: 5 + 1,
		},
		{
			StudentId:  3,
			QuestionId: 1,
		},
		{
			StudentId:  3,
			QuestionId: 2,
		},
		{
			StudentId:  3,
			QuestionId: 3,
		},
		{
			StudentId:  1,
			QuestionId: 5 + 1,
		},
		{
			StudentId:  1,
			QuestionId: 5 + 2,
		},
		{
			StudentId:  1,
			QuestionId: 5 + 3,
		},
		{
			StudentId:  1,
			QuestionId: 5 + 4,
		},
		{
			StudentId:  1,
			QuestionId: 5 + 5,
		},
		{
			StudentId:  1,
			QuestionId: 10 + 1,
		},
		{
			StudentId:  1,
			QuestionId: 10 + 2,
		},
		{
			StudentId:  1,
			QuestionId: 10 + 3,
		},
		{
			StudentId:  1,
			QuestionId: 10 + 4,
		},
		{
			StudentId:  1,
			QuestionId: 10 + 5,
		},
	}
	for _, machineStudent := range machineStudents {
		_, err := machineStudent.CreateQuestionStudent()
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func notificationTest() {
	notifications := []Notification{
		{
			Message: "New machine 'SQL Injection Basics' has been added to Web Security course",
			Type:    "announcement",
			SendAt:  time.Now(),
			Read:    false,
		},
		{
			Message: "Reminder: Complete the 'Buffer Overflow' challenge b",
			Type:    "reminder",
			SendAt:  time.Now(),
			Read:    true,
		},
		{
			Message: "System maintenance scheduled for Saturday, May 25th from 2-4 AM UTC",
			Type:    "system",
			SendAt:  time.Now(),
			Read:    false,
		},
		{
			Message: "Congratulations to the top 3 students on the leaderboard this month!",
			Type:    "announcement",
			SendAt:  time.Now(),
			Read:    true,
		},
		{
			Message: "New course 'Advanced Network Security' is now available",
			Type:    "reminder",
			SendAt:  time.Now(),
			Read:    false,
		},
	}

	for _, notificiation := range notifications {
		n, err := notificiation.CreateNotification()
		if err != nil {
			return
		}
		ni := &NotificationInstructor{
			InstructorId:   1,
			NotificationId: n.ID,
		}
		_, err = ni.CreateNotificationInstructor()
		if err != nil {
			return
		}
	}
	for i := 0; i < 100; i++ {
		notifiation := Notification{
			Message: generateRandomString(20),
			Type:    "system",
			SendAt:  time.Now(),
			Read:    false,
		}

		cn, err := notifiation.CreateNotification()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		ni := &NotificationInstructor{
			InstructorId:   1,
			NotificationId: cn.ID,
		}
		_, err = ni.CreateNotificationInstructor()
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
	notifications = []Notification{
		{
			Message: "New machine 'SQL Injection Basics' has been added to Web Security course",
			Type:    "announcement",
			SendAt:  time.Now(),
			Read:    false,
		},
		{
			Message: "Reminder: Complete the 'Buffer Overflow' challenge b",
			Type:    "reminder",
			SendAt:  time.Now(),
			Read:    true,
		},
		{
			Message: "System maintenance scheduled for Saturday, May 25th from 2-4 AM UTC",
			Type:    "system",
			SendAt:  time.Now(),
			Read:    false,
		},
		{
			Message: "Congratulations to the top 3 students on the leaderboard this month!",
			Type:    "announcement",
			SendAt:  time.Now(),
			Read:    true,
		},
		{
			Message: "New course 'Advanced Network Security' is now available",
			Type:    "reminder",
			SendAt:  time.Now(),
			Read:    false,
		},
	}

	for _, notificiation := range notifications {
		n, err := notificiation.CreateNotification()
		if err != nil {
			return
		}
		ni := &NotificationAdmin{
			AdminId:        1,
			NotificationId: n.ID,
		}
		_, err = ni.CreateNotificationAdmin()
		if err != nil {
			return
		}
	}

	for i := 0; i < 100; i++ {
		notifiation := Notification{
			Message: generateRandomString(20),
			Type:    "system",
			SendAt:  time.Now(),
			Read:    false,
		}

		cn, err := notifiation.CreateNotification()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		notifAd := &NotificationAdmin{
			AdminId:        1,
			NotificationId: cn.ID,
		}
		_, err = notifAd.CreateNotificationAdmin()
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}

	notifications = []Notification{
		{
			Message: "New machine 'SQL Injection Basics' has been added to Web Security course",
			Type:    "announcement",
			SendAt:  time.Now(),
			Read:    false,
		},
		{
			Message: "Reminder: Complete the 'Buffer Overflow' challenge b",
			Type:    "reminder",
			SendAt:  time.Now(),
			Read:    true,
		},
		{
			Message: "System maintenance scheduled for Saturday, May 25th from 2-4 AM UTC",
			Type:    "system",
			SendAt:  time.Now(),
			Read:    false,
		},
		{
			Message: "Congratulations to the top 3 students on the leaderboard this month!",
			Type:    "announcement",
			SendAt:  time.Now(),
			Read:    true,
		},
		{
			Message: "New course 'Advanced Network Security' is now available",
			Type:    "reminder",
			SendAt:  time.Now(),
			Read:    false,
		},
	}

	for _, notificiation := range notifications {
		n, err := notificiation.CreateNotification()
		if err != nil {
			return
		}

		ni := &NotificationStudent{
			StudentId:      1,
			NotificationId: n.ID,
		}
		_, err = ni.CreateNotificationStudent()
		if err != nil {
			return
		}
	}

}

func generateRandomString(length int) string {
	const (
		firstCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		otherCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)

	if length < 1 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Start with a valid first character
	result := make([]byte, length)
	result[0] = firstCharset[r.Intn(len(firstCharset))]

	// Fill the rest with valid subsequent characters
	for i := 1; i < length; i++ {
		result[i] = otherCharset[r.Intn(len(otherCharset))]
	}

	return string(result)
}

func PullImage() {
	dm, err := docker.NewDockerManager()
	if err != nil {
		log.Println(err.Error())
		return
	}
	images := []string{
		"mintesnotafework/sql_injection-image:v1.0",
		"mintesnotafework/xss_playground-image:v1.0",
		"mintesnotafework/command-injection-image:v1.0",
		"mintesnotafework/broken-auth-image:v1.0",
		"mintesnotafework/file-inclusion-image:v1.0",
		"mintesnotafework/sensitive_data_exposure-image:v1.0",
		"mintesnotafework/insecure-deserialization-image:v1.0",
		"mintesnotafework/browser-attackbox:v1.0",
	}
	for _, image := range images {
		err := dm.PullImage(context.Background(), image)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
