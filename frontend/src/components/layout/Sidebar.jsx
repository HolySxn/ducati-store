import React from 'react';
import { Bike, Package, Users, ShoppingCart, Settings, Tags } from 'lucide-react';
import '../../styles/Sidebar.css';

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
    <div className="sidebar">
      <div className="sidebar-header">
        <h1 className="sidebar-title">Ducati Admin</h1>
      </div>
      <nav className="sidebar-nav">
        {navItems.map((item) => (
          <button
            key={item.section}
            onClick={() => onSectionChange(item.section)}
            className={`nav-item ${activeSection === item.section ? 'active' : ''}`}
          >
            <item.icon className="nav-icon" />
            {item.name}
          </button>
        ))}
      </nav>
    </div>
  );
};

export default Sidebar;