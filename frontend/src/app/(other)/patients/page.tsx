'use client';

import { useEffect, useState } from 'react';
import axios from 'axios';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';

interface Patient {
  id: number;
  full_name: string;
  age: number;
  gender: 'male' | 'female' | 'other';
  address: string;
  phone: string;
  status: 'admitted' | 'discharged' | 'under observation';
}

const ListPatients = () => {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const router = useRouter();

  useEffect(() => {
    const fetchPatients = async () => {
      try {
        const response = await axios.get<Patient[]>('http://localhost:8080/patient/get-all', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('access_token')}`,
          },
        });

        setPatients(response.data);
      } catch (error: any) {
        console.error(error);
        toast.error(error.response?.data?.details || 'Failed to fetch patients');
      } finally {
        setLoading(false);
      }
    };

    fetchPatients();
  }, []);

  const handleDelete = async (id: number) => {
    try {
      const response = await axios.delete(`http://localhost:8080/patient/delete/${id}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('access_token')}`,
        },
      });
      toast.success('Patient deleted successfully');
      setPatients((prevPatients) => prevPatients.filter((patient) => patient.id !== id));
    } catch (error: any) {
      console.error(error);
      toast.error(error.response?.data?.details || 'Failed to delete patient');
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="min-h-screen bg-muted px-4">
      <h1 className="text-2xl font-bold text-center mb-6">Patients List</h1>

      <table className="table-auto w-full bg-white rounded shadow-md">
        <thead>
          <tr>
            <th className="border px-4 py-2">ID</th>
            <th className="border px-4 py-2">Full Name</th>
            <th className="border px-4 py-2">Age</th>
            <th className="border px-4 py-2">Gender</th>
            <th className="border px-4 py-2">Status</th>
            <th className="border px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {patients.map((patient) => (
            <tr key={patient.id}>
              <td className="border px-4 py-2">{patient.id}</td>
              <td className="border px-4 py-2">{patient.full_name}</td>
              <td className="border px-4 py-2">{patient.age}</td>
              <td className="border px-4 py-2">{patient.gender}</td>
              <td className="border px-4 py-2">{patient.status}</td>
              <td className="border px-4 py-2">
                <button
                  className="bg-yellow-500 text-white px-2 py-1 w-full rounded mb-2"
                  onClick={() => {
                    router.push(`/patients/${patient.id}`);
                  }}
                >
                  Edit
                </button>
                <button
                  className="bg-red-500 text-white px-2 py-1 w-full rounded"
                  onClick={() => handleDelete(patient.id)}
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ListPatients;
