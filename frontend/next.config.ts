import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
    experimental:{
        turbo:{
            treeShaking:true,
            memoryLimit: 1024 * 1024 * 512 // in bytes / 512MB
        }
    }
};

export default nextConfig;
