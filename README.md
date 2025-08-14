## Nourish

#### Notice

This project was developed as a learning experience to explore the full software development lifecycle, from conception to deployment. Due to the high cost associated with maintaining a live demo using the Spoonacular API, this application is effectively read-only and is not intended for public use or continuous maintenance.

![logo](/images/logo.png)

**URL:** _not public as the spoonacular api is rather expensive_

## Overview

Nourish is a full-featured meal planning web application designed to streamline the meal preparation process. Users can plan their meals, manage grocery lists, and save recipes based on dietary preferences and allergies. With an intuitive drag-and-drop dashboard, this application provides a seamless experience for managing weekly meals and grocery needs.

#### Features

- **Weekly Meal Planning:** Organize meals across the week with a drag-and-drop interface.
- **Recipe Management:** Save, view, and organize favorite recipes.
- **Grocery Lists:** Automatically generate grocery lists based on planned meals.
- **Dietary Preferences & Allergies:** Customize meal plans based on dietary needs and food intolerances.
- **Nutritional Information:** View nutritional details on each recipe card for easy access.
- **User Authentication:** Secure login with Google OAuth.

## How it works

**Deployment**

![image.png](/images/infra.png)

deployment is done via a digital ocean droplet. The reason for deploying to a VPS instead of any serverless solutions is that I gain much more control over the stack. To me, serverless feels like too much magic, and that magic runs the risk of inconsistent pricing. While yes, running this app serverless would likely be dirt cheap, what happens at a much larges scale? What happens if my API Gateway gets hit with a DDoS attack? I run the risk of paying an exuberant price for all those network calls.

**Authentication**

![image.png](/images/auth.png)

In past projects, I've had tons of auth troubles. Specifically, I've had troubles with regards to security. Because of this, I've realized that using third-party auth such as Auth0 and Clerk, is typically a better solution for public facing applications. **However, **using third party libraries doesn't help me grow as an engineer and it is for this reason, I've decided to handle a large majority of it manually.

**Database**

![image.png](/images/schema.png)

There are two primary approaches for building a meal planning application:

1. Users create recipes and share those recipes with others, manually
2. Recipes get pulled from an external API
   I chose the ladder option as I wanted to work with external data as I haven't had much experience doing so. Because of this, much of the apps functionality can be handled the API (Spoonacular).

#### Technology Stack

- **Frontend:** Nuxt with Tailwind CSS for styling.
- **Backend Services: **Golang and the standard library
- **Database:** PostgreSQL, with tables for user profiles, dietary preferences, and recipes.
- **CI/CD Pipeline:** GitHub Actions used for automated testing and deployment.
- **Deployment:** Docker Compose with Traefik as a reverse proxy, hosted on a DigitalOcean VPS with Cloudflare for added security.
- **CDN:** DigitalOcean Spaces for media assets.
- **Cache Management:** Nitro caching + redis

#### Future Improvements

- **Mobile Application:** Potential future expansion to a mobile app.
- **Additional Authentication Options:** Support for other OAuth providers.

#### License

This project is licensed under the MIT License. See the LICENSE file for more details.
