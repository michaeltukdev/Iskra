'use client'
import Link from "next/link"
import { Server, LayoutGrid, Cpu, CalendarDays, Settings, LifeBuoy, Bell, LucideIcon, ChevronRight, X, Menu } from 'lucide-react';
import { useState, memo } from "react";
import clsx from "clsx";
import { useAuth } from "@/context/AuthContext";

const navigationLinks = [
    { href: '/dashboard', label: 'Overview', icon: LayoutGrid },
    { href: '/dashboard/nodes', label: 'Node List', icon: Server },
    { href: '/dashboard/hardware', label: 'Hardware', icon: Cpu },
    { href: '/dashboard/calendar', label: 'Calendar', icon: CalendarDays },
    { href: '/dashboard/settings', label: 'Settings', icon: Settings },
]

interface NavLinkProps extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
    href: string
    label: string
    icon: LucideIcon
    className?: string
}

const NavLink: React.FC<NavLinkProps> = ({ href = '/', label = '', className = '', icon: Icon = LayoutGrid }) => {
    return (
        <Link href={href} className={clsx("flex items-center gap-3 px-2 py-2 text-sm font-normal text-text-medium transition-colors duration-200 hover:text-text-light", className)}>
            {Icon && <Icon size={16} />}
            {label}
        </Link>
    );
};

const SidebarHeader: React.FC = memo(() => {
    const { user } = useAuth();

    return (
        // TODO: Update with user profile link
        <Link href="/" className="flex items-center justify-between p-3.5 max-h-[70px] border-b border-foreground-border text-sm gap-2 group">
            <div className="flex items-center gap-2">
                <div className="w-10 h-10 bg-foreground-border rounded-lg"></div>
                <div>
                    <p>{user?.username}</p>
                    <p className="text-xs text-text-spare font-medium">{user?.email}</p>
                </div>
            </div>
            <ChevronRight size={16} className="group-hover:text-text-light transition-all" />
        </Link>
    )
});

const SidebarContent: React.FC = memo(() => (
    <>
        <SidebarHeader />
        <div className="p-4 flex-grow overflow-y-auto">
            <input type="text" placeholder="Search..." className="w-full bg-transparent border border-input rounded-lg px-2.5 py-2 text-text-medium text-sm focus:outline-none focus:border-primary hover:border-primary transition-all mb-4 focus:ring-0" />
            <NavLink href="/" label="Alerts" icon={Bell} />

            <p className="text-sm text-text-dark mt-4 mb-2">Navigation</p>

            {navigationLinks.map((link, index) => (
                <NavLink key={index} href={link.href} label={link.label} icon={link.icon} />
            ))}

        </div>
        <div className="p-4">
            {/* TODO: Update with documentation link */}
            <NavLink href="/" label="Documentation" icon={LifeBuoy} />
        </div>
    </>
));

const Sidebar: React.FC = () => {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <>
            <div className="bg-foreground w-full p-3.5 lg:hidden flex justify-between items-center">
                <span className="text-lg font-semibold">Navigation</span>
                <button onClick={() => setIsOpen(!isOpen)} className="text-text-medium hover:text-text-light transition-all focus:outline-none" aria-label="Toggle Sidebar" >
                    {isOpen ? <X size={24} /> : <Menu size={24} />}
                </button>
            </div>

            {isOpen && (
                <>
                    <div onClick={() => setIsOpen(false)} className="fixed lg:hidden inset-0 bg-black bg-opacity-30 z-40" aria-hidden="true" ></div>

                    <div className={`fixed top-0 left-0 w-full max-w-[280px] h-full bg-foreground flex flex-col z-50 transform transition-transform duration-300 ${isOpen ? 'translate-x-0' : '-translate-x-full'} lg:hidden`} aria-label="Sidebar" >
                        <SidebarContent />
                    </div>
                </>
            )}

            <div className="hidden lg:flex bg-foreground w-[280px] h-screen flex-col">
                <SidebarContent />
            </div>
        </>
    );
}

export default Sidebar;