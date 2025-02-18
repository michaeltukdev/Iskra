'use client';

import { Input } from '@/components/UI/Input';
import { useState } from 'react';
import Link from 'next/link';
import { validateRegister } from '@/utils/validation/authValidation';
import { useAuth } from '@/context/AuthContext';
import { useRouter } from 'next/navigation';

export default function Register() {
  const { register } = useAuth();
  const router = useRouter();
  const [username, setUsername] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [passwordConfirmation, setPasswordConfirmation] = useState<string>('');
  const [localError, setLocalError] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    setLocalError(null);

    if (password !== passwordConfirmation) {
      setLocalError({ passwordConfirmation: 'Passwords do not match' });
      return;
    }

    const validationErrors = validateRegister({ username, email, password });
    if (validationErrors) {
      setLocalError(validationErrors);
      return;
    }

    setLoading(true);
    // TODO: When registration open toggle exists, redirect to login page (will also add middleware, but just in case)
    
    try {
      await register(username, email, password);
      router.push('/auth/login');
    } catch (backendErrors) {
      setLocalError(backendErrors);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <form onSubmit={handleSubmit} className="w-full max-w-md bg-foreground p-5 rounded-md">
        <div className='space-y-5'>
          <Input required placeholder="Username" label="Username" type="text" value={username} onChange={(e) => setUsername(e.target.value)} error={localError?.username} />

          <Input required placeholder="Enter your email..." label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} error={localError?.email} />

          <Input showToggle required placeholder="Enter your password..." label="Password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} error={localError?.password} />

          <Input showToggle required placeholder="Confirm your password..." label="Password Confirmation" type="password" value={passwordConfirmation} onChange={(e) => setPasswordConfirmation(e.target.value)} error={localError?.passwordConfirmation} />
        </div>

        {localError?.error && <p className="text-red-400 text-sm mt-2">{localError.error}</p>}

        <button type="submit" className="w-full bg-primary hover:bg-primary-hover transition-all text-white py-2.5 rounded-md text-sm mt-6">
          {loading ? "Creating Account..." : "Register"}
        </button>
      </form>

      <p className='mt-5 text-sm'>Have an account? <Link className='text-text-light hover:text-primary transition-all' href="/auth/login">Sign In</Link></p>
    </div>
  )
}