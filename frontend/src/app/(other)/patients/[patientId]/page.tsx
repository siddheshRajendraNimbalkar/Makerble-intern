'use client';

import { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams, useRouter } from 'next/navigation';
import { toast } from 'sonner';

interface Patient {
  id: number;
  full_name: string;
  age: number;
  gender: 'male' | 'female' | 'other';
  address: string;
  phone: string;
  status: 'admitted' | 'discharged' | 'under observation';
}

const PatientDetails = () => {
  const params = useParams();
  const patientId = params.patientId as string; // Cast patientId as string

  const [patient, setPatient] = useState<Patient | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [isUpdating, setIsUpdating] = useState<boolean>(false);

  const [formData, setFormData] = useState({
    full_name: '',
    age: 0,
    gender: 'male',
    address: '',
    phone: '',
    status: 'admitted',
  });

  const router = useRouter();

  useEffect(() => {
    if (!patientId) {
      toast.error('Invalid patient ID');
      return;
    }

    const fetchPatientDetails = async () => {
      try {

        const response = await axios.get<Patient>(`http://localhost:8080/patient/get/${patientId}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('access_token')}`,
          },
        });

        setPatient(response.data);
        setFormData({
          full_name: response.data.full_name,
          age: response.data.age,
          gender: response.data.gender,
          address: response.data.address.String,
          phone: response.data.phone.String,
          status: response.data.status,
        });
      } catch (error: any) {
        console.error(error);
        toast.error('Failed to fetch patient details');
      } finally {
        setLoading(false);
      }
    };

    fetchPatientDetails();
  }, [patientId]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsUpdating(true);

    try {

      const updatedData = {
        ...formData,
        age: Number(formData.age), 
      };
      let res = await axios.put(`http://localhost:8080/patient/update/${patientId}`, updatedData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('access_token')}`,
        },
      });

      toast.success('Patient updated successfully');
      router.push('/patients'); 
    } catch (error: any) {
      toast.error('Failed to update patient');
    } finally {
      setIsUpdating(false);
    }
  };

  if (loading) {
    return <div>Loading patient details...</div>;
  }

  if (!patient) {
    return <div>No patient found</div>;
  }

  return (
    <div className="min-h-screen bg-muted px-4">
      <h1 className="text-2xl font-bold text-center mb-6">Patient Details</h1>

      <form onSubmit={handleSubmit} className="space-y-4 max-w-lg mx-auto">
        <div className="flex flex-col">
          <label className="mb-2" htmlFor="full_name">
            Full Name
          </label>
          <input
            type="text"
            name="full_name"
            id="full_name"
            value={formData.full_name}
            onChange={handleChange}
            className="px-3 py-2 border rounded-md"
          />
        </div>

        <div className="flex flex-col">
          <label className="mb-2" htmlFor="age">
            Age
          </label>
          <input
            type="number"
            name="age"
            id="age"
            value={formData.age}
            onChange={handleChange}
            className="px-3 py-2 border rounded-md"
          />
        </div>

        <div className="flex flex-col">
          <label className="mb-2" htmlFor="gender">
            Gender
          </label>
          <select
            name="gender"
            id="gender"
            value={formData.gender}
            onChange={handleChange}
            className="px-3 py-2 border rounded-md"
          >
            <option value="male">Male</option>
            <option value="female">Female</option>
            <option value="other">Other</option>
          </select>
        </div>

        <div className="flex flex-col">
          <label className="mb-2" htmlFor="address">
            Address
          </label>
          <input
            type="text"
            name="address"
            id="address"
            value={formData.address}
            onChange={handleChange}
            className="px-3 py-2 border rounded-md"
          />
        </div>

        <div className="flex flex-col">
          <label className="mb-2" htmlFor="phone">
            Phone
          </label>
          <input
            type="text"
            name="phone"
            id="phone"
            value={formData.phone}
            onChange={handleChange}
            className="px-3 py-2 border rounded-md"
          />
        </div>

        <div className="flex flex-col">
          <label className="mb-2" htmlFor="status">
            Status
          </label>
          <select
            name="status"
            id="status"
            value={formData.status}
            onChange={handleChange}
            className="px-3 py-2 border rounded-md"
          >
            <option value="admitted">Admitted</option>
            <option value="discharged">Discharged</option>
            <option value="under observation">Under Observation</option>
          </select>
        </div>

        <div className="flex justify-center">
          <button
            type="submit"
            className="bg-blue-500 text-white px-4 py-2 rounded-md"
            disabled={isUpdating}
          >
            {isUpdating ? 'Updating...' : 'Update Patient'}
          </button>
        </div>
      </form>
    </div>
  );
};

export default PatientDetails;
