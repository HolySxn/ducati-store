import React, { useState, useEffect } from 'react';
import { AlertCircle } from 'lucide-react';
import { Alert, AlertDescription } from '@/components/ui/alert';
import Sidebar from '../components/layout/Sidebar';
import DataTable from '../components/tables/DataTable';
import { fetchData, deleteItem } from '../services/api';
import MotorcycleList from './Motorcycles/MotorcycleList';
// import CategoryList from './Categories/CategoryList';
// import CustomerList from './Customers/CustomerList';
// import OrderList from './Orders/OrderList';
// import ServiceRecordList from './ServiceRecords/ServiceRecordList';
// import InventoryList from './Inventory/InventoryList';
import '../styles/Dashboard.css';

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
      case 'categories':
        return <CategoryList />;
      case 'customers':
        return <CustomerList />;
      case 'orders':
        return <OrderList />;
      case 'service':
        return <ServiceRecordList />;
      case 'inventory':
        return <InventoryList />;
      default:
        return <div>Section under construction</div>;
    }
  };

  return (
    <div className="dashboard-container">
      <Sidebar activeSection={activeSection} onSectionChange={setActiveSection} />
      <div className="main-content">
        {error && (
          <Alert variant="destructive" className="mb-4">
            <AlertCircle />
            <AlertDescription>{error}</AlertDescription>
          </Alert>
        )}

        {loading ? (
          <div className="flex justify-center items-center h-full">
            <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"></div>
          </div>
        ) : (
          <div className="space-y-6">
            <div className="section-header">
              <h2>{activeSection.charAt(0).toUpperCase() + activeSection.slice(1)}</h2>
              <button className="add-button">Add New</button>
            </div>
            
            {renderSection()}

            <div className="table-container">
              <DataTable
                data={data}
                columns={Object.keys(data[0] || {}).filter(key => key !== '_id')}
                onEdit={(item) => {/* TODO: Edit modal */}}
                onDelete={handleDelete}
              />
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Dashboard;
