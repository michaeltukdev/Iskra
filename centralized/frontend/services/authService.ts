const API_URL = process.env.API_URL;

export const authService = {
    register: async (email: string, username: string, password: string) => {
        const response = await fetch(API_URL + "/auth/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, username, password }),
        });

        if (!response.ok) {
            const data = await response.json();
            throw new Error(data.message || "Validation failed");
        }

        return response.json();
    },

    login: async (email: string, password: string) => {
        try {
            const response = await fetch(API_URL + "/auth/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ email, password }),
                credentials: "include",
            });

            if (!response.ok) {
                const data = await response.json();
                throw new Error(data.message || "An unexpected error occurred.");
            }

            return await response.json();
        } catch (error: any) {
            throw new Error(error.message || "Unable to connect to the server. Please try again.");
        }
    },

    logout: async () => {
        const response = await fetch(API_URL + "/auth/logout", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
        });

        if (!response.ok) {
            const data = await response.json();
            throw new Error(data.message || "Logout failed");
        }

        return response.json();
    },

    validate: async () => {
        try {
            const response = await fetch(API_URL + "/me", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            });

            if (!response.ok) {
                if (response.status === 401) {
                    return null;
                }

                const errorData = await response.json();
                throw new Error(errorData.message || "Validation failed");
            }

            const data = await response.json();
            return data.Claims;
        } catch (error) {
            throw error;
        }
    },
}