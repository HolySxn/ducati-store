import React, { useState, useEffect } from 'react';
import { AlertCircle, Plus, Loader2 } from 'lucide-react';
import { Alert, AlertDescription } from '@/components/ui/alert';
import DataTable from '../../components/tables/DataTable';
import MotorcycleForm from './MotorcycleForm';
import { fetchData, deleteItem } from '../../services/api';
import '../../styles/MotorcycleList.css';

const MotorcycleList = () => {
  const [motorcycles, setMotorcycles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [editingMotorcycle, setEditingMotorcycle] = useState(null);
  const [sortConfig, setSortConfig] = useState({ key: null, direction: 'asc' });

  const columns = [
    {
      key: 'Model',
      header: 'Model',
      sortable: true
    },
    {
      key: 'Brand',
      header: 'Brand',
      sortable: true
    },
    {
      key: 'Category',
      header: 'Category',
      sortable: true
    },
    {
      key: 'Year',
      header: 'Year',
      sortable: true
    },
    {
      key: 'Price',
      header: 'Price',
      sortable: true,
      render: (value) => `$${value?.toLocaleString() || 0}`
    },
    {
      key: 'Status',
      header: 'Status',
      sortable: true,
      render: (value) => (
        <span className={`status-badge ${value?.toLowerCase()}`}>
          {value || 'N/A'}
        </span>
      )
    },
    {
      key: 'Specifications',
      header: 'Specifications',
      render: (specs) => {
        if (!specs) return 'N/A';
        return (
          <div className="specs-container">
            {specs.map((spec, index) => (
              <span key={index} className="spec-item">
                {spec}
              </span>
            ))}
          </div>
        );
      }
    }
  ];

  const fetchMotorcycles = async () => {
    setLoading(true);
    try {
      const motorcycles = await fetchData('motorcycles');
      const categories = await fetchData('categories');
      console.log('Fetched motorcycles:', motorcycles);
      console.log('Fetched categories:', categories);

      const motorcyclesWithCategories = motorcycles.map(motorcycle => {
        const category = categories.find(cat => cat.ID === motorcycle.CategoryID);
        return { ...motorcycle, Category: category ? category.Name : 'Unknown' };
      });

      console.log('Motorcycles with categories:', motorcyclesWithCategories);
      setMotorcycles(motorcyclesWithCategories);
      setError(null);
    } catch (err) {
      setError('Failed to fetch motorcycles. Please try again later.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchMotorcycles();
  }, []);

  const handleDelete = async (id) => {
    if (!window.confirm('Are you sure you want to delete this motorcycle?')) {
      return;
    }

    try {
      await deleteItem('motorcycles', id);
      await fetchMotorcycles();
    } catch (err) {
      setError('Failed to delete motorcycle. Please try again.');
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

  const handleSort = (key) => {
    setSortConfig((prevSort) => ({
      key,
      direction:
        prevSort.key === key && prevSort.direction === 'asc' ? 'desc' : 'asc',
    }));
  };

  const sortedMotorcycles = React.useMemo(() => {
    if (!sortConfig.key) return motorcycles;

    const sorted = [...motorcycles].sort((a, b) => {
      if (a[sortConfig.key] < b[sortConfig.key]) {
        return sortConfig.direction === 'asc' ? -1 : 1;
      }
      if (a[sortConfig.key] > b[sortConfig.key]) {
        return sortConfig.direction === 'asc' ? 1 : -1;
      }
      return 0;
    });

    console.log('Sorted motorcycles:', sorted);
    return sorted;
  }, [motorcycles, sortConfig]);

  return (
    <div className="motorcycle-list">
      {error && (
        <Alert variant="destructive" className="error-alert">
          <AlertCircle className="alert-icon" />
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}

      <div className="list-header">
        <h1 className="page-title">Motorcycles</h1>
        <button
          onClick={() => setIsFormOpen(true)}
          className="add-button"
        >
          <Plus className="button-icon" />
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
        <div className="loading-container">
          <Loader2 className="loading-spinner" />
          <span>Loading motorcycles...</span>
        </div>
      ) : (
        <DataTable
          data={sortedMotorcycles}
          columns={columns}
          onEdit={handleEdit}
          onDelete={handleDelete}
          onSort={handleSort}
          sortConfig={sortConfig}
        />
      )}
    </div>
  );
};

export default MotorcycleList;