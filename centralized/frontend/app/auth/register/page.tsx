'use client';
import { useState } from "react";
import { authService } from "@/services/authService";
import { splitError } from "@/utils/helpers";
import { ErrorObject } from "@/app/types/utils";
import { useRouter } from 'next/navigation'

export default function Register() {
  const router = useRouter();
  
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errors, setErrors] = useState<ErrorObject>();
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    setErrors({});
    setIsLoading(true);

    try {
      const response = await authService.register(email, username, password);
      setIsLoading(false);

      router.push("/login");   
    } catch (error) {
      setIsLoading(false);
      if (error instanceof Error) {
        const errors = splitError(error.message)
        setErrors(errors);
      }
    }
  };

  return (
    <div>
      <form className="flex-col space-y-8" onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email</label>
          <input type="email" id="email" name="email" value={email} onChange={(e) => setEmail(e.target.value)} />
          {errors && errors.email && <p>{errors.email}</p>}
        </div>

        <div>
          <label htmlFor="username">Username</label>
          <input type="text" id="username" name="username" value={username} onChange={(e) => setUsername(e.target.value)} />
          {errors && errors.username && <p>{errors.username}</p>}
        </div>

        <div>
          <label htmlFor="password">Password</label>
          <input type="password" id="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
          {errors && errors.password && <p>{errors.password}</p>}

        </div>

        <button type="submit" disabled={isLoading}>
          {isLoading ? "Registering..." : "Register"}
        </button>

        {errors && errors.error && <p>{errors.error}</p>}

      </form>

    </div>
  );
}
