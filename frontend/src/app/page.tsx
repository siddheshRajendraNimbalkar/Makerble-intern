'use client';
import { Button } from "@/components/ui/button";
import Image from "next/image";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();
  return (
    <>
    HELLO
    <br />
    <Button variant="outline" className="bg-black text-white"
      onClick={()=> {
        router.push('/login')
      }}
    >Login</Button>
    <Button variant="outline" className="bg-black text-white"
      onClick={()=>{
        router.push('/register')
      }}
    >register</Button>
    <br />
    <Button variant="outline" onClick={() => {
      router.push('/receptionist')
    }}>Create Patient</Button>
    <Button variant="outline" onClick={() => {
      router.push('/patients')
    }}>list Patient</Button>
    
    </>
  );
}
