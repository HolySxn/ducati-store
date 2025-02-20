import React from 'react';
import { Pencil, Trash2 } from 'lucide-react';

const DataTable = ({ data, columns, onEdit, onDelete, onSort, sortConfig = {} }) => {
  if (!data || data.length === 0) {
    return <div className="empty-state">No data available</div>;
  }
  
  return (
    <div className="table-container">
      <table className="data-table">
        <thead>
          <tr>
            {columns.map((column) => (
              <th 
                key={column.key}
                onClick={() => column.sortable && onSort(column.key)}
                className={column.sortable ? 'sortable' : ''}
              >
                {column.header}
                {sortConfig.key === column.key && (
                  <span className="sort-indicator">
                    {sortConfig.direction === 'asc' ? ' ↑' : ' ↓'}
                  </span>
                )}
              </th>
            ))}
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {data.map((item, index) => (
            <tr key={item.ID || index}>
              {columns.map((column) => (
                <td key={`${item.ID || index}-${column.key}`}>
                  {column.render ? column.render(item[column.key]) : item[column.key]}
                </td>
              ))}
              <td className="actions">
                <button 
                  onClick={() => onEdit(item)} 
                  className="action-button edit"
                >
                  <Pencil size={16} />
                </button>
                <button 
                  onClick={() => onDelete(item.ID)} 
                  className="action-button delete"
                >
                  <Trash2 size={16} />
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default DataTable;