package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dawamr/resume-cv-2023/pkg/env"
	"github.com/gofiber/fiber/v2"
)

func RenderHello(c *fiber.Ctx) (err error) {
	experience := getExp();
	skills := getSkill();
	portfolio := getPortofolio();

	url := "https://cms.dawam.my.id/items/Blog?sort=-date_created&fields=*,category_id.name,banner.*"
	resp, err := http.Get(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error fetching data")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error reading response body")
	}

	var bodyBlog APIResponse
	if err := json.Unmarshal(body, &bodyBlog); err != nil {
		fmt.Println(body)
		bodyBlog = APIResponse{}
	}
	blogs := bodyBlog.Data
	
	return c.Render("index", fiber.Map{
		"CMS_HOST":        env.GetEnv("CMS_HOST", "https://cms.dawam.my.id"),
		"Name":            "Dawam Raja",
		"Email":           "id.dawamraja@gmail.com",
		"Phone":           "+6285174427553",
		"FiberTitle":      "Resume Dawam Raja",
		"JobPositions":    []string{"Software Engineer", "Project Manager", "Backend Developer", "Coder"},
		"Address":         []string{"Jakarta", "Indonesia"},
		"Company":         []string{"Solarion", "PT Solarion Energi Alam"},
		"StartExperience": 2017,
		"ExperienceYears": calculateExperienceYears(2017),
		"ExperienceJob":   experience,
		"ExperienceSkill": skills,
		"ExperiencePortofolio": portfolio,
		"Blog" : blogs,
	})
}

type Experience struct {
	Image       string
	Date        string
	Job         string
	Location    string
	ShortDesc   string
	Description string
}

type Skill struct {
	Image       string
	Title       string
	Description string
	Details     []string
}

type PortfolioItem struct {
	Image       string
	Category    string
	Title       string
	Description string
	Details     []string
}

type ContactData struct {
    Name    string
    Message string
}

// Category represents the structure of a category in the API response
type Category struct {
	Category string `json:"name"`
}

type Banner struct {
	ID           string    `json:"id"`
	Storage      string    `json:"storage"`
	FilenameDisk string    `json:"filename_disk"`
	FilenameDownload string `json:"filename_download"`
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	Folder       string    `json:"folder"`
	UploadedBy   string    `json:"uploaded_by"`
	UploadedOn   time.Time `json:"uploaded_on"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedOn   time.Time `json:"modified_on"`
	Filesize     string    `json:"filesize"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
}


type Blog struct {
	ID           int        `json:"id"`
	Status       string     `json:"status"`
	UserCreated  string     `json:"user_created"`
	DateCreated  time.Time  `json:"date_created"`
	UserUpdated  string     `json:"user_updated"`
	DateUpdated  time.Time  `json:"date_updated"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	Category     Category    `json:"category_id"`
	Banner       Banner     `json:"banner"`
	Tag	  		[]string	`json:"tag"`
}

type APIResponse struct {
	Data []Blog `json:"data"`
}

func calculateExperienceYears(startExperience int) int {
	return time.Now().Year() - startExperience
}

func ContactSend(c *fiber.Ctx) error {
	contact := ContactData{
		Name:    c.FormValue("name"),
		Message: c.FormValue("message"),
	}

	 // Generate WhatsApp URL
        phone := "6285174427553" // Replace with your WhatsApp phone number
        whatsappURL := fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s",
            url.QueryEscape(phone),
            url.QueryEscape(fmt.Sprintf("Name: %s\nMessage: %s", contact.Name, contact.Message)),
        )

		return c.Redirect(whatsappURL, fiber.StatusSeeOther)
}

func getPortofolio() []PortfolioItem{
	return []PortfolioItem{
		{
			Image:       "/img/portfolio/coworking.jpeg",
			Category:    "KodingWorks",
			Title:       "Coworking & Doorlock App",
			Description: "Coworking & Doorlock App - An integrated application tailored for efficient coworking space management, allowing users to book coworking spaces, order additional facilities, and seamlessly integrate with doorlock systems.",
			Details: []string{
				"This app streamlines the coworking experience, offering features such as flexible space reservations and the ability to request and manage additional facilities.",
				"Built using a dynamic technology stack, including Flutter & Swift for engaging cross-platform mobile interfaces, Node.js & Laravel for robust backend operations, and MySQL for secure and efficient data management.",
			},
		},
		{
			Image:       "/img/portfolio/hotel.jpeg",
			Category:    "KodingWorks",
			Title:       "Hotel Booking App",
			Description: "Hotel Booking App - A specialized application tailored for seamless hotel reservations, room, food, and souvenir orders at one of the hotels in Indonesia.",
			Details: []string{
				"This app simplifies the booking process, allowing users to reserve rooms, browse food catalogs, order souvenirs, make secure payments, and monitor their bookings.",
				"Built using a versatile technology stack, including React Native & Swift for engaging mobile interfaces, NestJS & Laravel for robust backend operations, and MySQL database for secure and efficient data management.",
			},
		},
		{
			Image:       "/img/portfolio/tumbasin.jpeg",
			Category:    "KodingWorks",
			Title:       "Shop Mart & Delivery App",
			Description: "Shop Mart & Delivery App - An innovative application designed for a seamless shopping experience at traditional markets across several cities.",
			Details: []string{
				"This app facilitates shopping, order scheduling, driver-assisted delivery, and secure payment options for customers. Additionally, it provides insightful analytics on customer activities, enhancing shopping experiences.",
				"Built using a technology stack that includes ReactJS for dynamic user interfaces, Express & Laravel for efficient backend operations, and MySQL database for secure and organized data management.",
			},
		},
		{
			Image:       "/img/portfolio/machinevision.jpeg",
			Category:    "KodingWorks",
			Title:       "Machine Vision Monitoring App",
			Description: "Machine Vision Monitoring App - A cutting-edge application designed to monitor various aspects within a manufacturing facility, including plants, machines, production lines, and specific sectors or machine components.",
			Details: []string{
				"This application employs a sophisticated calculation matrix, facilitating error detection, work monitoring, and anomaly detection. The use of this matrix allows for precise and timely notifications regarding the status of machines and processes, enhancing overall operational efficiency.",
				"Built using a technology stack that includes ReactJS for the dynamic user interface, Express for seamless backend functionality, and MongoDB for efficient data management and storage.",
			},
		},
		{
			Image:       "/img/portfolio/klinis.jpeg",
			Category:    "Terretech",
			Title:       "Telemedicine & Klinik App",
			Description: "Telemedicine & Klinik App - An integrated application designed to provide telemedicine services, enabling users to engage in chat and video calls with medical professionals, order medications and medical tools, schedule vaccinations, make clinic reservations, and utilize AI-powered medical checks.",
			Details: []string{
				"This app also facilitates payments through multiple methods, including DANA, various payment gateways, and Moota. It seamlessly connects with essential services such as WhatsApp, AWS Cognito for authentication, and OAuth2 for secure authorization.",
				"Built using a robust technology stack, including React Native & Next.js for dynamic mobile interfaces, Laravel & Node.js for efficient backend operations, and MySQL & Redis for secure and efficient data management.",
			},
		},
	}
}

func getSkill() []Skill  {
	return []Skill{
		{
			Image:       "/img/service/2.jpg",
			Title:       "Backend Engineer",
			Description: "As a Backend Engineer proficient in Golang, Python, Node.js, and PHP, I specialize in developing robust and scalable server-side solutions to power modern web applications.",
			Details: []string{
				"Golang (Go): Golang is a powerful, statically typed programming language known for its simplicity, efficiency, and concurrency support. I leverage Golang to create high-performance and concurrent backend solutions, particularly for microservices architecture.",
				"Python: Python is a versatile, high-level programming language known for its readability and vast ecosystem of libraries and frameworks. I use Python for various backend tasks, such as data processing, automation, and web development.",
				// Add more details...
			},
		},
		{
			Image:       "/img/service/3.jpg",
			Title:       "Frontend Developer",
			Description: "As a Frontend Developer specializing in React, Vue.js, and Laravel, I excel in crafting seamless and interactive user interfaces by seamlessly integrating with backends, utilizing RESTful APIs, and providing robust debugger support.",
			Details: []string{
				"React: React is a powerful JavaScript library for building user interfaces. I leverage React's component-based structure and virtual DOM to develop interactive and responsive frontend applications, providing a delightful user experience.",
				"Vue.js: Vue is a progressive JavaScript framework used for building modern, reactive web interfaces. I utilize Vue.js to create efficient and maintainable frontend solutions, allowing for seamless integration with backend services.",
				// Add more details...
			},
		},
		{
			Image:       "/img/service/4.jpg",
			Title:       "Database Management",
			Description: "As a database specialist, I have extensive experience in managing various database systems, including MySQL, PostgreSQL, MongoDB, Redis, and Elasticsearch, ensuring efficient data storage, retrieval, and optimization.",
			Details: []string{
				"MySQL: MySQL is a popular open-source relational database management system known for its reliability, performance, and ease of use. It is widely used for various applications, from small-scale projects to enterprise-level solutions.",
				"PostgreSQL: PostgreSQL is a powerful, open-source, object-relational database system known for its advanced features, extensibility, and strong community support. It is often chosen for applications that require complex queries and transactions.",
				// Add more database details...
			},
		},
		{
			Image:       "/img/service/5.jpg",
			Title:       "Project Manager",
			Description: "As a Project Manager, I possess the skills and expertise required to oversee and successfully manage diverse projects from initiation to completion.",
			Details: []string{
				"Project Planning and Organization: Skillful in developing comprehensive project plans, setting objectives, defining roles, and allocating resources efficiently to meet project goals within defined timelines.",
				"Team Leadership and Collaboration: Proficient in leading cross-functional teams, fostering collaboration, and promoting a positive working environment to achieve project milestones and deliverables.",
				// Add more project management details...
			},
		},
		{
			Image:       "/img/service/6.jpg",
			Title:       "Team Leader",
			Description: "As a Team Lead, I excel in guiding and motivating teams to achieve collective success while fostering individual growth and collaboration.",
			Details: []string{
				"Leadership and Guidance: Proficient in providing clear direction, setting goals, and fostering a cohesive team environment to maximize productivity and achieve project objectives.",
				"Team Development: Skilled in identifying strengths and weaknesses within the team, and implementing strategies for professional growth, skill enhancement, and career advancement.",
				// Add more team leader details...
			},
		},
	}
}
func getExp() []Experience {
	return []Experience{
		{
			Image:       "/img/experience/1.jpg",
			Date:        "Oct 2020 - Present",
			Job:         "IT Project Manager",
			Location:    "Solarion Energi Alam, Jakarta",
			ShortDesc:   "As an IT Project Manager at PT Solarion Energi Alam, I lead and oversee various projects, ensuring their successful delivery and alignment with business goals.",
			Description: "PT Solarion Energi Alam is the holding company of PT Terretech Nusantara, PT Kreasi Layanan Media, and several other sub-holdings. In my role, I focus on analytics, technical engineering, project management, and leadership to drive project success and business growth.",
		},
		{
			Image:       "/img/experience/2.jpg",
			Date:        "Jul 2019 - Feb 2020",
			Job:         "Back End Developer",
			Location:    "Koding Kreasi Indonesia, Semarang",
			ShortDesc:   "During my time at Koding Kreasi Indonesia, I was a dedicated Back End Developer involved in website development, programming, coding, and maintenance of web applications.",
			Description: "Koding Kreasi Indonesia focused on intensive coding bootcamps for 'yatim' and 'dhuafa,' emphasizing training, book writing, and utilizing technologies such as Laravel and JavaScript libraries.",
		},
		{
			Image:       "/img/experience/3.jpg",
			Date:        "Feb 2018 - Jun 2019",
			Job:         "Full Stack Engineer",
			Location:    "Santren Koding, Semarang",
			ShortDesc:   "At Santren Koding, I served as a Full Stack Engineer, contributing to the development, programming, and maintenance of websites and web applications.",
			Description: "Santren Koding, a foundation focusing on coding bootcamps for 'yatim' and 'dhuafa,' honed my skills in training, book writing, and technologies like Laravel and JavaScript libraries.",
		},
		{
			Image:       "/img/experience/4.jpg",
			Date:        "Jan 2018 - Mar 2018",
			Job:         "Intern",
			Location:    "Santren Koding, Semarang",
			ShortDesc:   "During this period, I was engaged as an intern, expanding my knowledge and skills in the field of software development.",
			Description: "I dedicated this time to enhance my understanding of software development and related technologies, preparing myself for future endeavors.",
		},
	}
}
