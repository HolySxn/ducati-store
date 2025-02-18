import React from 'react';
import { Bike, Package, Users, ShoppingCart, Settings, Tags } from 'lucide-react';

const navItems = [
  { name: 'Motorcycles', icon: Bike, section: 'motorcycles' },
  { name: 'Categories', icon: Tags, section: 'categories' },
  { name: 'Inventory', icon: Package, section: 'inventory' },
  { name: 'Customers', icon: Users, section: 'customers' },
  { name: 'Orders', icon: ShoppingCart, section: 'orders' },
  { name: 'Service Records', icon: Settings, section: 'service' },
];

const Sidebar = ({ activeSection, onSectionChange }) => {
  return (
    <div className="w-64 bg-white shadow-lg">
      <div className="p-6">
        <h1 className="text-2xl font-bold text-gray-800">Ducati Admin</h1>
      </div>
      <nav className="mt-6">
        {navItems.map((item) => (
          <button
            key={item.section}
            onClick={() => onSectionChange(item.section)}
            className={`w-full flex items-center p-4 text-sm ${
              activeSection === item.section
                ? 'bg-indigo-500 text-white'
                : 'text-gray-600 hover:bg-gray-50'
            }`}
          >
            <item.icon className="w-5 h-5 mr-3" />
            {item.name}
          </button>
        ))}
      </nav>
    </div>
  );
};

export default Sidebar;