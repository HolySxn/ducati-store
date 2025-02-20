import React, { useState, useEffect } from 'react';
import { AlertCircle } from 'lucide-react';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { createItem, updateItem, fetchData } from '../../services/api';
import '../../styles/MotorcycleForm.css';

const MotorcycleForm = ({ motorcycle = null, onSave, onCancel }) => {
  const [formData, setFormData] = useState({
    model: '',
    brand: '',
    year: new Date().getFullYear(),
    price: 0,
    specifications: [],
    status: 'Available',
    categoryId: ''
  });
  const [categories, setCategories] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const categoriesData = await fetchData('categories');
        setCategories(categoriesData);
      } catch (err) {
        setError('Failed to fetch categories');
      }
    };

    fetchCategories();
  }, []);

  useEffect(() => {
    if (motorcycle) {
      setFormData({
        model: motorcycle.Model || '',
        brand: motorcycle.Brand || '',
        year: motorcycle.Year || new Date().getFullYear(),
        price: motorcycle.Price || 0,
        specifications: motorcycle.Specifications || [],
        status: motorcycle.Status || 'Available',
        categoryId: motorcycle.CategoryID || ''
      });
    }
  }, [motorcycle]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (motorcycle) {
        await updateItem('motorcycles', motorcycle.ID, formData);
      } else {
        await createItem('motorcycles', formData);
      }
      onSave();
    } catch (err) {
      setError('Failed to save motorcycle');
    }
  };

  const handleSpecificationChange = (index, value) => {
    const newSpecs = [...formData.specifications];
    newSpecs[index] = value;
    setFormData({ ...formData, specifications: newSpecs });
  };

  const addSpecification = () => {
    setFormData({
      ...formData,
      specifications: [...formData.specifications, '']
    });
  };

  const removeSpecification = (index) => {
    const newSpecs = formData.specifications.filter((_, i) => i !== index);
    setFormData({ ...formData, specifications: newSpecs });
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow form-container" >
      <h2 className="text-xl font-bold mb-4 form-title">
        {motorcycle ? 'Edit Motorcycle' : 'Add New Motorcycle'}
      </h2>
      
      {error && (
        <Alert variant="destructive" className="mb-4 alert alert-error">
          <AlertCircle className="h-4 w-4" />
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}

      <form onSubmit={handleSubmit} className="space-y-4 form-content">
        <div className='form-group'>
          <label className="block text-sm font-medium text-gray-700 form-label">Model</label>
          <input
            type="text"
            value={formData.model}
            onChange={(e) => setFormData({ ...formData, model: e.target.value })}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 form-input"
            required
          />
        </div>

        <div className='form-group'>
          <label className="block text-sm font-medium text-gray-700 form-label">Brand</label>
          <input
            type="text"
            value={formData.brand}
            onChange={(e) => setFormData({ ...formData, brand: e.target.value })}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 form-input"
            required
          />
        </div>

        <div className='form-group'>
          <label className="block text-sm font-medium text-gray-700 form-label">Year</label>
          <input
            type="number"
            value={formData.year}
            onChange={(e) => setFormData({ ...formData, year: parseInt(e.target.value) })}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 form-input"
            required
          />
        </div>

        <div className='form-group'>
          <label className="block text-sm font-medium text-gray-700 form-label">Price</label>
          <input
            type="number"
            value={formData.price}
            onChange={(e) => setFormData({ ...formData, price: parseFloat(e.target.value) })}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 form-input"
            required
          />
        </div>

        <div className='form-group'>
          <label className="block text-sm font-medium text-gray-700 form-label">Status</label>
          <select
            value={formData.status}
            onChange={(e) => setFormData({ ...formData, status: e.target.value })}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 form-input"
          >
            <option value="Available">Available</option>
            <option value="Not Available">Not Available</option>
          </select>
        </div>

        <div className='form-group'>
          <label className="block text-sm font-medium text-gray-700 form-label">Category</label>
          <select
            value={formData.categoryId}
            onChange={(e) => setFormData({ ...formData, categoryId: e.target.value })}
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 form-input"
            required
          >
            <option value="">Select a category</option>
            {categories.map((category) => (
              <option key={category.ID} value={category.ID}>
                {category.Name}
              </option>
            ))}
          </select>
        </div>

        <div className='form-group'>
          <label className="block text-sm font-medium text-gray-700 mb-2 form-label">
            Specifications
          </label>
          {formData.specifications.map((spec, index) => (
            <div key={index} className="flex gap-2 mb-2">
              <input
                type="text"
                value={spec}
                onChange={(e) => handleSpecificationChange(index, e.target.value)}
                className="flex-1 rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 form-input"
                placeholder="e.g., 1000cc, 150hp, etc."
              />
              <button
                type="button"
                onClick={() => removeSpecification(index)}
                className="px-2 py-1 text-red-600 hover:text-red-800 button"
              >
                Remove
              </button>
            </div>
          ))}
          <button
            type="button"
            onClick={addSpecification}
            className="text-indigo-600 hover:text-indigo-800 text-sm button"
          >
            + Add Specification
          </button>
        </div>

        <div className="flex gap-4 button-group">
          <button
            type="submit"
            className="px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600 button button-primary"
          >
            {motorcycle ? 'Update' : 'Create'} Motorcycle
          </button>
          <button
            type="button"
            onClick={onCancel}
            className="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 button button-secondary"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
};

export default MotorcycleForm;