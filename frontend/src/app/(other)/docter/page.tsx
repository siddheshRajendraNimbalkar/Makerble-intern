'use client';

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

const docter = () => {
    const router = useRouter();
    return (
        <div className="min-h-screen flex items-center justify-center bg-muted px-4">
        <h1 className="text-2xl font-bold text-center">Docter</h1>
        <br />
        <Button variant="outline" onClick={() => {
            router.push('/patients')
        }}>list Patient</Button>
        </div>
    );
}

export default docter;