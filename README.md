## Nourish

**Project Name:** Help Me Meal Prep
**URL:** *not public yet*

#### Overview

Nourish is a full-featured meal planning web application designed to streamline the meal preparation process. Users can plan their meals, manage grocery lists, and save recipes based on dietary preferences and allergies. With an intuitive drag-and-drop dashboard, this application provides a seamless experience for managing weekly meals and grocery needs.

#### Features

- Weekly Meal Planning: Organize meals across the week with a drag-and-drop interface.
- Recipe Management: Save, view, and organize favorite recipes.
- Grocery Lists: Automatically generate grocery lists based on planned meals.
- Dietary Preferences & Allergies: Customize meal plans based on dietary needs and food intolerances.
- Nutritional Information: View nutritional details on each recipe card for easy access.
- User Authentication: Secure login with Google OAuth.

#### Technology Stack

- **Frontend:** Nuxt.js with Tailwind CSS for styling.
- **Backend:** Nuxt Nitro server, structured as a monolith but with some functionality separated for clarity.
- **Database:** PostgreSQL, with tables for user profiles, dietary preferences, and recipes.
- **CI/CD Pipeline:** GitHub Actions used for automated testing and deployment.
- **Deployment:** Docker Compose with Traefik as a reverse proxy, hosted on a DigitalOcean VPS with Cloudflare for added security.
- **CDN:** DigitalOcean Spaces for media assets.
- **Cache Management:** Manual caching implemented in some services to optimize recipe information retrieval.

#### Project Structure

- **Frontend:** Nuxt.js project with Tailwind CSS.
- **Backend:** Nuxt Nitro API server with structured services (e.g., recipeService).
- **CI/CD:** GitHub Actions for testing (using Vitest) and deployment pipeline with GitHub Packages.

#### Future Improvements

- **Mobile Application:** Potential future expansion to a mobile app.
- **Enhanced Caching:** Better management of cached data for improved performance.
- **Additional Authentication Options:** Support for other OAuth providers.

#### License

This project is licensed under the MIT License. See the LICENSE file for more details.
