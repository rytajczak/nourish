<p><a target="_blank" href="https://app.eraser.io/workspace/WO6FPCcUiJGt5l8xozHX" id="edit-in-eraser-github-link"><img alt="Edit in Eraser" src="https://firebasestorage.googleapis.com/v0/b/second-petal-295822.appspot.com/o/images%2Fgithub%2FOpen%20in%20Eraser.svg?alt=media&amp;token=968381c8-a7e7-472a-8ed6-4a6626da5501"></a></p>

## Capstone
![logo](/public/logo.png "")

**URL:** _not public yet_

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

![image.png](/.eraser/WO6FPCcUiJGt5l8xozHX___1Enb9iXhRwP7bhK676Cd8i3xr483___Dce2iWTcfmT4pTyrg4eZC.png "image.png")

deployment is done via a digital ocean droplet (because I dislike serverless). The actual reason for deploying to a VPC instead of any serverless solutions is that I gain much more control over the stack. To me, serverless feels like too much magic, and that magic runs the risk of inconsistent pricing. While yes, running this app serverless would likely be dirt cheap, what happens at a much larges scale? What happens if my API Gateway gets hit with a DDoS attack? I run the risk of paying an exuberant price for all those network calls. 

**Authentication**

![image.png](/.eraser/WO6FPCcUiJGt5l8xozHX___1Enb9iXhRwP7bhK676Cd8i3xr483___HPzssizIGiXVdES9Aeov6.png "image.png")

In past projects, I've had tons of auth troubles. Specifically, I've had troubles with regards to security. Because of this, I've realized that using third-party auth such as Auth0 and Clerk, is typically a better solution for public facing applications. **However, **using third party libraries doesn't help me grow as an engineer and it is for this reason, I've decided to handle a large majority of it manually.

**Database**

![image.png](/.eraser/WO6FPCcUiJGt5l8xozHX___1Enb9iXhRwP7bhK676Cd8i3xr483___pMXUTSoWJHqlhI0ER3QjC.png "image.png")

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


<!-- eraser-additional-content -->
## Diagrams
<!-- eraser-additional-files -->
<a href="/README-sequence-diagram-1.eraserdiagram" data-element-id="OUojz27MvAF2S1eSahmlM"><img src="/.eraser/WO6FPCcUiJGt5l8xozHX___1Enb9iXhRwP7bhK676Cd8i3xr483___---diagram----f6cf10d8e3687933d452c976251eb7a8.png" alt="" data-element-id="OUojz27MvAF2S1eSahmlM" /></a>
<a href="/README-entity-relationship-2.eraserdiagram" data-element-id="05pVl2cFFD3UAyDDaEHR0"><img src="/.eraser/WO6FPCcUiJGt5l8xozHX___1Enb9iXhRwP7bhK676Cd8i3xr483___---diagram----68cbb7fba1f2af347162699b9dfe2a7f.png" alt="" data-element-id="05pVl2cFFD3UAyDDaEHR0" /></a>
<!-- end-eraser-additional-files -->
<!-- end-eraser-additional-content -->
<!--- Eraser file: https://app.eraser.io/workspace/WO6FPCcUiJGt5l8xozHX --->