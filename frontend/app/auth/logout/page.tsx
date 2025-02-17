'use client'
import { authService } from "@/services/authService";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Logout() {
    const router = useRouter();

    useEffect(() => {
        authService.logout()
            .then(() => {
                router.push("/auth/login");
            })
            .catch((error) => {
                console.error("Logout failed:", error);
            });
    }, [router]);

    return null;
}
