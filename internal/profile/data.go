package profile

import models "github.com/pr4th4meshh/p3rsonal_API/internal/model"

func ProfileData() models.Profile {
	return models.Profile{
		Name:    "Prathamesh Asolkar",
		Handle:  "pr4th4meshh",
		Title:   "Full-Stack Developer",
		Email:   "prathameshasolkar***@gmail.com",
		Phone:   "+91 750609xxxx",
		Summary: "Passionate Full-Stack Developer with 2+ years of hands-on experience building responsive, user-friendly websites and web applications with modern, in-demand technologies.",
		Links: models.Links{
			Website:  "https://iamprathamesh.xyz",
			LinkedIn: "https://www.linkedin.com/in/prathamesh-asolkar/",
			GitHub:   "https://github.com/pr4th4meshh",
		},
		Experience: []models.Experience{
			{
				Slug:     "prana-india",
				Role:     "SDE-1",
				Company:  "Prana India",
				Location: "Bengaluru (On-site)",
				Start:    "Sept 2025",
				End:      "Present",
				Stack:    []string{"Next.js", "React", "TypeScript", "Zustand", "Zod", "TanStack Query", "WebSockets"},
				Highlights: []string{
					"Led development of the customer-facing pranaindia.com using Next.js and TypeScript, enabling the company's transition from offline to online.",
					"Served 400K+ weekly active users and supported 1.2M+ monthly revenue, with 4.8M+ monthly revenue being the highest.",
					"Contributed to 22.85M+ cumulative revenue from Jan-Jul 2026.",
					"Engineered real-time chat and in-app voice consultations end-to-end using WebSockets, replacing paid third-party APIs with an in-house solution for secure, low-latency communication.",
					"Built the complete consultation platform using React, TypeScript, Zustand, Zod, and TanStack Query, delivering consultation workflows, push notifications, analytics dashboards, and scalable real-time user experiences.",
				},
			},
			{
				Slug:     "orim-advisors",
				Role:     "Software Developer Intern",
				Company:  "ORIM Advisors",
				Location: "Mumbai (On-site)",
				Start:    "Jan 2025",
				End:      "Sept 2025",
				Stack:    []string{"Vite", "React", "TypeScript", "TanStack Query", "Zustand", "Docker", "Playwright", "Vitest", "FastAPI", "Python"},
				Highlights: []string{
					"Built a scalable, high-performance front-end for cloud software using Vite, React, TypeScript, TanStack Query, Zustand, and Docker in collaboration with Cloud Engineers.",
					"Led development and maintenance of alphaco.ai, contributing to product growth and supporting its successful investment raise.",
					"Delivered robust testing through E2E and component tests with Playwright and Vitest, built FastAPI routes, and contributed to Python data-ingestion pipelines.",
				},
			},
			{
				Slug:     "fuzzynodes",
				Role:     "Full Stack Developer Intern",
				Company:  "Fuzzynodes",
				Location: "Australia (Remote)",
				Start:    "Dec 2024",
				End:      "June 2025",
				Stack:    []string{"Next.js", "Node.js", "Express", "TypeScript", "Bun", "Elysia", "Docker", "GitHub Actions", "AWS Amplify", "OpenAI GPT-4"},
				Highlights: []string{
					"Engineered and deployed high-performance web applications: a multi-restaurant platform and admin dashboard using Next.js, Node.js, Express, TypeScript, Bun, and Elysia.",
					"Made 20+ REST API endpoints.",
					"Set up CI/CD pipelines using Docker, GitHub Actions, and AWS Amplify, ensuring a smooth and reliable deployment architecture.",
					"Integrated OpenAI GPT-4 for backend functions including intelligent profanity detection and automated user-blacklisting auto-fill.",
				},
			},
		},
		Projects: []models.Project{
			{
				Slug:         "skribbl-clone",
				Name:         "Skribbl Clone",
				Subtitle:     "Real-time Multiplayer Drawing Game",
				Live:         true,
				Technologies: []string{"React", "TypeScript", "Node.js", "Socket.io", "Redis", "BullMQ", "JWT", "Zustand", "TanStack Query", "HTML5 Canvas"},
				Highlights: []string{
					"Built a real-time multiplayer drawing-and-guessing game.",
					"Developed turn-based gameplay with Redis-backed room state and BullMQ queues for timers, word reveals, and round management across concurrent sessions.",
					"Implemented JWT authentication with refresh-token rotation and efficient client/server state management with Zustand and TanStack Query.",
				},
			},
			{
				Slug:         "presssence",
				Name:         "Presssence",
				Subtitle:     "AI Portfolio Builder",
				Live:         true,
				Technologies: []string{"Next.js", "TypeScript", "MongoDB", "Prisma", "Gemini API", "RAG", "Vector Embeddings", "Streaming"},
				Highlights: []string{
					"Built a full-stack portfolio builder with a 5-step onboarding flow, shareable portfolios, inline editing, and drag-and-drop reordering.",
					"Shipped a resume-to-portfolio importer using schema-constrained LLM output.",
					"Built a public RAG chatbot with vector embeddings and streaming responses.",
					"Built a provider-agnostic AI layer with rate limiting, retries, and an eval harness with a 100% pass rate.",
				},
			},
		},
		Education: []models.Education{
			{
				Institution: "Bhavan's College",
				Affiliation: "University of Mumbai",
				Degree:      "Better don't ask :P",
				Start:       "Aug 2021",
				End:         "May 2024",
			},
		},
		Skills: models.Skills{
			Languages: []string{"JavaScript", "TypeScript", "Python", "Go"},
			Databases: []string{"MongoDB", "PostgreSQL", "Redis"},
			Frontend:  []string{"React", "Next.js", "Tailwind", "Redux", "Zustand", "TanStack"},
			Backend:   []string{"Node.js", "Express", "FastAPI", "Bun", "Elysia", "Prisma"},
			Mobile:    []string{"React Native", "Expo", "Nativewind"},
			Testing:   []string{"Vitest", "Playwright", "PWAs"},
			DevOps:    []string{"Docker", "AWS (EC2, S3, ECR, Amplify)"},
			Tools:     []string{"Git", "Insomnia", "Postman", "Jira", "Slack", "Miro", "LLMs"},
		},
		Rules: []string{
			"Source-of-truth is in-process Go data; no database is required.",
			"Search is case-insensitive and matches names, roles, companies, subtitles, highlights, technologies, and skills.",
			"Unknown resource slugs return HTTP 404 with a structured error.",
			"No external API is called by the server.",
			"Resume-listed URLs are never invented; missing link URLs stay empty.",
			"The API is read-only by design: there is no generic CRUD endpoint for the personal profile.",
		},
	}
}
