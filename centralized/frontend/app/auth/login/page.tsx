'use client';

import { Input } from '@/components/UI/Input';
import { useState } from 'react';
import { useAuth } from '@/context/AuthContext';
import { validateLogin } from '@/utils/validation/authValidation';
import { useRouter } from 'next/navigation'
import Link from 'next/link';

export default function Login() {
    const { login } = useAuth();
    const router = useRouter();
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [localError, setLocalError] = useState<any>(null);
    const [loading, setLoading] = useState<boolean>(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        setLocalError(null);
        const validationErrors = validateLogin({ email, password });

        if (validationErrors) {
            setLocalError(validationErrors);
            return;
        }

        setLoading(true);

        try {
            await login(email, password);
            router.push('/');
        } catch (backendErrors) {
            setLocalError(backendErrors);
        } finally {
            setLoading(false);
        }
    };


    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <form onSubmit={handleSubmit} className="w-full max-w-md bg-foreground p-5 rounded-md">
                <div className='space-y-5'>
                    <Input required placeholder="Enter your email..." label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} error={localError?.email} />

                    <Input required showToggle placeholder="Enter your password..." label="Password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} error={localError?.password} />
                </div>

                {localError?.error && <p className="text-red-400 text-sm mt-2">{localError.error}</p>}

                <button type="submit" className="w-full bg-primary hover:bg-primary-hover transition-all text-white py-2.5 rounded-md text-sm mt-6">
                    {loading ? "Signing In" : "Sign In"}
                </button>
            </form>

            {/* TODO: When registration open toggle exists, toggle this */}
            <p className='mt-5 text-sm'>Not registered with us? <Link className='text-text-light hover:text-primary transition-all' href="/auth/register">Create a new account</Link></p>
        </div>
    );
}
