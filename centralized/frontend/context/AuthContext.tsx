'use client';

import { authService } from '@/services/authService';
import React, { createContext, useState, useEffect, ReactNode, useContext } from 'react';
import { User } from '@/types/global';
import { splitError } from '@/utils/helpers';

interface AuthContextType {
  isAuthenticated: boolean;
  user: User | null;
  register: (username: string, email: string, password: string) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  isLoading: boolean;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const user = await authService.validate();

        if (user) {
          setIsAuthenticated(true);
          setUser(user);
        }

      } catch (error) {
        // throw error;
      } finally {
        setIsLoading(false);
      }
    }

    fetchUser();
  }, []);

  const register = async (username: string, email: string, password: string) => {
    try {
      const response = await authService.register(email, username, password);

      if (response.message) {
        throw new Error(response.message);
      }

      setIsAuthenticated(true);
      setUser(response.user);
    } catch (error: any) {
      const structuredError = splitError(error.message);
      throw structuredError;
    }
  }

  const login = async (email: string, password: string) => {
    try {
      const response = await authService.login(email, password);

      if (response.message) {
        throw new Error(response.message);
      }

      setIsAuthenticated(true);
      setUser(response.user);
    } catch (error: any) {
      const structuredError = splitError(error.message);
      throw structuredError;
    }
  };

  const logout = async () => {

  }

  return (
    <AuthContext.Provider value={{ isAuthenticated, user, register, login, logout, isLoading }} >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);

  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }

  return context;
};
