import type { NextConfig } from "next";
const path = require('path');
const dotenv = require('dotenv');

dotenv.config({ path: path.resolve(__dirname, '../.env') });

const nextConfig: NextConfig = {
  env: {
    API_URL: process.env.BACKEND_URL
  }
};

export default nextConfig;
