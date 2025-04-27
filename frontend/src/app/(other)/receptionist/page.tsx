'use client';

import { useForm } from 'react-hook-form';
import { useRouter } from 'next/navigation';
import axios from 'axios';
import { toast } from 'sonner';

type PatientFormValues = {
  full_name: string;
  age: number; // age as number
  gender: 'male' | 'female' | 'other';
  address: string;
  phone: string;
  status: 'admitted' | 'discharged' | 'under observation';
};

const ReceptionistPage = () => {
  const { register, handleSubmit, reset } = useForm<PatientFormValues>();
  const router = useRouter();

  const onSubmit = async (data: PatientFormValues) => {
    try {
      console.log('Form data:', data);
      await axios.post('http://localhost:8080/patient/create', data, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('access_token')}`,
        },
      });
      toast.success('Patient created successfully!');
      reset();
      router.refresh();
    } catch (error: any) {
      console.error(error);
      toast.error(error.response?.data?.details || 'Something went wrong');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-muted px-4">
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-4 bg-white p-6 rounded shadow-md w-full max-w-md">
        <h1 className="text-2xl font-bold text-center mb-6">Create Patient</h1>

        <input
          type="text"
          {...register('full_name', { required: 'Full name is required' })}
          placeholder="Full Name"
          className="w-full border px-3 py-2 rounded text-sm"
        />

        <input
          type="number"
          {...register('age', { required: 'Age is required', min: { value: 0, message: 'Invalid age' }, valueAsNumber: true })}
          placeholder="Age"
          className="w-full border px-3 py-2 rounded text-sm"
        />

        <select
          {...register('gender', { required: 'Gender is required' })}
          className="w-full border px-3 py-2 rounded text-sm"
        >
          <option value="">Select Gender</option>
          <option value="male">Male</option>
          <option value="female">Female</option>
          <option value="other">Other</option>
        </select>

        <input
          type="text"
          {...register('address', { required: 'Address is required' })}
          placeholder="Address"
          className="w-full border px-3 py-2 rounded text-sm"
        />

        <input
          type="text"
          {...register('phone', { required: 'Phone number is required' })}
          placeholder="Phone Number"
          className="w-full border px-3 py-2 rounded text-sm"
        />

        <select
          {...register('status', { required: 'Status is required' })}
          className="w-full border px-3 py-2 rounded text-sm"
        >
          <option value="">Select Status</option>
          <option value="admitted">Admitted</option>
          <option value="discharged">Discharged</option>
          <option value="under observation">Under Observation</option>
        </select>

        <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition">
          Create Patient
        </button>
      </form>
    </div>
  );
};

export default ReceptionistPage;
