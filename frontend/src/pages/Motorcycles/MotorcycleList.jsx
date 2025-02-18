import React, { useState, useEffect } from 'react';
import { AlertCircle } from 'lucide-react';
import { Alert, AlertDescription } from '@/components/ui/alert';
import DataTable from '../../components/tables/DataTable';
import MotorcycleForm from './MotorcycleForm';
import { fetchData, deleteItem } from '../../services/api';

const MotorcycleList = () => {
  const [motorcycles, setMotorcycles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [editingMotorcycle, setEditingMotorcycle] = useState(null);

  const fetchMotorcycles = async () => {
    setLoading(true);
    try {
      const data = await fetchData('motorcycles');
      setMotorcycles(data);
      setError(null);
    } catch (err) {
      setError('Failed to fetch motorcycles');
    }
    setLoading(false);
  };

  useEffect(() => {
    fetchMotorcycles();
  }, []);

  const handleDelete = async (id) => {
    try {
      await deleteItem('motorcycles', id);
      fetchMotorcycles();
    } catch (err) {
      setError('Failed to delete motorcycle');
    }
  };

  const handleEdit = (motorcycle) => {
    setEditingMotorcycle(motorcycle);
    setIsFormOpen(true);
  };

  const handleSave = async () => {
    setIsFormOpen(false);
    setEditingMotorcycle(null);
    await fetchMotorcycles();
  };

  const tableColumns = [
    'model',
    'brand',
    'year',
    'price',
    'status',
    'specifications'
  ];

  return (
    <div className="p-6">
      {error && (
        <Alert variant="destructive" className="mb-4">
          <AlertCircle className="h-4 w-4" />
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}

      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold text-gray-800">Motorcycles</h1>
        <button
          onClick={() => setIsFormOpen(true)}
          className="px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600"
        >
          Add New Motorcycle
        </button>
      </div>

      {isFormOpen ? (
        <MotorcycleForm
          motorcycle={editingMotorcycle}
          onSave={handleSave}
          onCancel={() => {
            setIsFormOpen(false);
            setEditingMotorcycle(null);
          }}
        />
      ) : loading ? (
        <div className="flex justify-center items-center h-32">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
        </div>
      ) : (
        <DataTable
          data={motorcycles}
          columns={tableColumns}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      )}
    </div>
  );
};

export default MotorcycleList;