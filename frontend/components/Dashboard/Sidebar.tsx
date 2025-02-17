'use client';
import Link from "next/link";
import { Server, LayoutGrid, Cpu, CalendarDays, Settings, LifeBuoy, Bell, ChevronRight, X, Menu, ChevronDown, LucideIcon } from 'lucide-react';
import { useState, memo } from "react";
import clsx from "clsx";
import { useAuth } from "@/context/AuthContext";

const navigationLinks = [
    { href: '/dashboard', label: 'Overview', icon: LayoutGrid },
    {
        href: '/dashboard/nodes',
        label: 'Node List',
        icon: Server,
        children: [
            { href: '/dashboard/nodes', label: 'Node list' },
            { href: '/dashboard/nodes/create', label: 'Create Node' },
        ]
    },
    {
        href: '/dashboard/hardware',
        label: 'Hardware', icon: Cpu,
        children: [
            { href: '/dashboard/hardware/cpu', label: 'Hardware list' },
            { href: '/dashboard/hardware/cpu', label: 'Update Hardware' },
        ]
    },
    { href: '/dashboard/calendar', label: 'Calendar', icon: CalendarDays },
    {
        label: 'Settings',
        icon: Settings,
        children: [
            { href: '/dashboard/settings/profile', label: 'Profile' },
            { href: '/dashboard/settings/security', label: 'Security' },
        ],
    },
];

interface NavLinkProps extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
    href?: string;
    label: string;
    icon?: LucideIcon;
    className?: string;
    children?: never;
}

const NavLink: React.FC<NavLinkProps> = ({ href = '/', label, className = '', icon: Icon = '' }) => (
    <Link href={href} className={clsx("flex items-center gap-3 px-2 py-2 text-sm font-normal text-text-medium transition-colors duration-200 hover:text-text-light", className)}>
        {Icon && <Icon size={16} />}
        {label}
    </Link>
);

interface NavDropdownProps {
    label: string;
    icon: LucideIcon;
    children: { href: string; label: string }[];
}

const NavDropdown: React.FC<NavDropdownProps> = ({ label, icon: Icon, children }) => {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <div className="flex flex-col">
            <button
                onClick={() => setIsOpen(!isOpen)}
                className="flex items-center justify-between gap-3 px-2 py-2 text-sm font-normal text-text-medium transition-colors duration-200 hover:text-text-light">
                <div className="flex items-center gap-3">
                    {Icon && <Icon size={16} />}
                    {label}
                </div>
                <ChevronDown size={16} className={clsx("transition-transform", { 'rotate-180': isOpen })} />
            </button>
            {isOpen && (
                <div className="ml-4 pl-3 flex flex-col border-l border-input-border">
                    {children.map((child, index) => (
                        <NavLink key={index} href={child.href} label={child.label} />
                    ))}
                </div>
            )}
        </div>
    );
};

const SidebarHeader: React.FC = memo(() => {
    const { user } = useAuth();

    return (
        <Link href="/" className="flex items-center justify-between p-3.5 max-h-[70px] border-b border-foreground-border text-sm gap-2 group">
            <div className="flex items-center gap-2">
                <div className="w-10 h-10 bg-foreground-border rounded-lg"></div>
                <div>
                    <p>{user ? user?.username : "Loading..."}</p>
                    <p className="text-xs text-text-spare font-medium">{user ? user?.email : "Loading..."}</p>
                </div>
            </div>
            <ChevronRight size={16} className="group-hover:text-text-light transition-all" />
        </Link>
    );
});

const SidebarContent: React.FC = memo(() => (
    <>
        <SidebarHeader />
        <div className="p-4 flex-grow overflow-y-auto">
            <input type="text" placeholder="Search..." className="w-full bg-transparent border border-input rounded-lg px-2.5 py-2 text-text-medium text-sm focus:outline-none focus:border-primary hover:border-primary transition-all mb-4 focus:ring-0" />
            <NavLink href="/" label="Alerts" icon={Bell} />

            <p className="text-sm text-text-dark mt-4 mb-2">Navigation</p>

            {navigationLinks.map((link, index) => (
                link.children ? (
                    <NavDropdown key={index} label={link.label} icon={link.icon} children={link.children} />
                ) : (
                    <NavLink key={index} href={link.href} label={link.label} icon={link.icon} />
                )
            ))}
        </div>
        <div className="p-4">
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
                <button onClick={() => setIsOpen(!isOpen)} className="text-text-medium hover:text-text-light transition-all focus:outline-none" aria-label="Toggle Sidebar">
                    {isOpen ? <X size={24} /> : <Menu size={24} />}
                </button>
            </div>

            {isOpen && (
                <>
                    <div onClick={() => setIsOpen(false)} className="fixed lg:hidden inset-0 bg-black bg-opacity-30 z-40" aria-hidden="true"></div>

                    <div className={`fixed top-0 left-0 w-full max-w-[280px] h-full bg-foreground flex flex-col z-50 transform transition-transform duration-300 ${isOpen ? 'translate-x-0' : '-translate-x-full'} lg:hidden`} aria-label="Sidebar">
                        <SidebarContent />
                    </div>
                </>
            )}

            <div className="hidden lg:flex bg-foreground w-full max-w-[280px] h-screen flex-col">
                <SidebarContent />
            </div>
        </>
    );
};

export default Sidebar;
