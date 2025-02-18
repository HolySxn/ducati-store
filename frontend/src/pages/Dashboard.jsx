import React, { useState, useEffect } from 'react';
import { AlertCircle } from 'lucide-react';
import { Alert, AlertDescription } from '@/components/ui/alert';
import Sidebar from '../components/layout/Sidebar';
import DataTable from '../components/tables/DataTable';
import { fetchData, deleteItem } from '../services/api';
import MotorcycleList from './Motorcycles/MotorcycleList';
// Import other section components as they're created

const Dashboard = () => {
  const [activeSection, setActiveSection] = useState('motorcycles');
  const [data, setData] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    handleFetchData();
  }, [activeSection]);

  const handleFetchData = async () => {
    setLoading(true);
    try {
      const result = await fetchData(activeSection);
      setData(result);
      setError(null);
    } catch (err) {
      setError('Failed to fetch data');
    }
    setLoading(false);
  };

  const handleDelete = async (id) => {
    try {
      await deleteItem(activeSection, id);
      handleFetchData();
    } catch (err) {
      setError('Failed to delete item');
    }
  };

  const renderSection = () => {
    switch (activeSection) {
      case 'motorcycles':
        return <MotorcycleList />;
      // Add other sections as they're created
      default:
        return <div>Section under construction</div>;
    }
  };

  return (
    <div className="flex h-screen bg-gray-100">
      <Sidebar activeSection={activeSection} onSectionChange={setActiveSection} />
      <div className="flex-1 overflow-auto p-8">
        {error && (
          <Alert variant="destructive" className="mb-4">
            <AlertCircle className="h-4 w-4" />
            <AlertDescription>{error}</AlertDescription>
          </Alert>
        )}

        {loading ? (
          <div className="flex justify-center items-center h-full">
            <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
          </div>
        ) : (
          <div className="space-y-6">
            <div className="flex justify-between items-center">
              <h2 className="text-2xl font-bold text-gray-800">
                {activeSection.charAt(0).toUpperCase() + activeSection.slice(1)}
              </h2>
              <button
                onClick={() => {/* TODO: Add new item modal */}}
                className="px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600"
              >
                Add New
              </button>
            </div>
            
            {renderSection()}

            <DataTable
              data={data}
              columns={Object.keys(data[0] || {}).filter(key => key !== '_id')}
              onEdit={(item) => {/* TODO: Edit modal */}}
              onDelete={handleDelete}
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default Dashboard;