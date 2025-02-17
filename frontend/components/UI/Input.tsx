'use client'
import React from 'react';
import { useState } from 'react';
import { Eye, EyeOff } from 'lucide-react';

type InputProps = {
    label: string;
    type: string;
    value: string;
    placeholder: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    error?: string;
    showToggle?: boolean;
    required?: boolean;
};

const InputBaseStyles = `w-full rounded-md py-2.5 px-3.5 text-sm bg-input shadow-sm focus:border-input-border outline-none ring-0 focus:ring-0 transition-all`;

export const Input: React.FC<InputProps> = ({ label, type, value, placeholder, required = false, showToggle = false, onChange, error }) => {
    const inputClasses = `${InputBaseStyles} ${error ? 'border border-red-400' : 'border border-input-border'}`;

    const [isPasswordVisible, setIsPasswordVisible] = useState(false);
    const inputType = showToggle && type === 'password' ? (isPasswordVisible ? 'text' : 'password') : type;

    const togglePasswordVisibility = () => {
        setIsPasswordVisible(prev => !prev);
    };

    return (
        <div>
            <label className="block text-sm text-text-medium mb-2.5">{label}</label>

            <div className='relative'>
                <input required={required} placeholder={placeholder} autoComplete={type} type={inputType} value={value} onChange={onChange} className={inputClasses} />

                {showToggle && type === 'password' && (
                    <button type="button" onClick={togglePasswordVisibility} className="absolute right-3 top-1/2 transform -translate-y-1/2">
                        {isPasswordVisible ? <Eye size={20} /> : <EyeOff size={20} />}
                    </button>
                )}
            </div>

            {error && <p className="text-red-400 text-sm mt-2">{error}</p>}
        </div>
    );
};
