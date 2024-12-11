export const authService = {
    register: async (email: string, username: string, password: string) => {
        const response = await fetch("http://localhost:8080/auth/register", {
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

    login: async(email: string, password: string) => {
        const response = await fetch("http://localhost:8080/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
        });

        if (!response.ok) {
            const data = await response.json();
            throw new Error(data.message || "Validation failed");
        }

        return response.json();
    },

    logout: () => localStorage.removeItem("token"),

    getToken: () => localStorage.getItem("token"),

    saveToken: (token: string) => {
        localStorage.setItem("token", token);
    }
}