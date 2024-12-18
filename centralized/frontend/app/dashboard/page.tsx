import { useAuth } from "@/context/AuthContext"
import Link from "next/link"

export default function Dashboard() {
    const { isAuthenticated, user, isLoading } = useAuth()

    if (isLoading) {
        return <p>Loading...</p>
    }

    if (!isAuthenticated || !user) {
        return (
            <div>
                <h1>Home</h1>
                <p>Welcome to the home page</p>
                <Link href="/auth/login">Login</Link>
            </div>
        )
    }

    return (
        <div>
            <h1>Dashboard</h1>
        </div>
    )
}