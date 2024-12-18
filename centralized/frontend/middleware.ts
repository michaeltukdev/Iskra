import { NextRequest, NextResponse } from "next/server";

export function middleware(req: NextRequest) {
    const url = req.nextUrl;

    if (url.pathname === "/auth/login" || url.pathname === "/auth/register") {
        const token = req.cookies.get("token")?.value;

        if (token) {
            return NextResponse.redirect(new URL("/", req.url));
        }
    }

    if (url.pathname === "/dashboard") {
        const token = req.cookies.get("token")?.value;

        if (!token) {
            return NextResponse.redirect(new URL("/auth/login", req.url));
        }
    }

    return NextResponse.next();
}

export const config = {
    matcher: ["/auth/:path*", "/dashboard/:path*"],
}