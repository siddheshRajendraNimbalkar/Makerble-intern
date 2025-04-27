'use client';

import React from 'react';
import { useTokenChecker } from '../tokenChecker';

const Layout = ({ children }: { children: React.ReactNode }) => {
  useTokenChecker(); 

  return (
    <html lang="en">
      <body className="bg-muted">
        {children}
      </body>
    </html>
  );
}

export default Layout;
