import Sidebar from "@/components/Dashboard/Sidebar"

export default function DashboardLayout({ children }: Readonly<{ children: React.ReactNode }>) {
    return (
        <div className="lg:flex">
            <Sidebar />

            <div className="lg:px-[70px]">
                {children}
            </div>
        </div>
    )
}