import React from 'react';

const DataTable = ({ data, columns, onEdit, onDelete }) => (
  <div className="overflow-x-auto bg-white rounded-lg shadow">
    <table className="min-w-full">
      <thead className="bg-gray-50">
        <tr>
          {columns.map(column => (
            <th key={column} className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              {column}
            </th>
          ))}
          <th className="px-6 py-3">Actions</th>
        </tr>
      </thead>
      <tbody className="divide-y divide-gray-200">
        {data.map((item) => (
          <tr key={item._id}>
            {columns.map(column => (
              <td key={column} className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {typeof item[column] === 'object' ? JSON.stringify(item[column]) : item[column]}
              </td>
            ))}
            <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <button onClick={() => onEdit(item)} className="text-indigo-600 hover:text-indigo-900 mr-4">
                Edit
              </button>
              <button onClick={() => onDelete(item._id)} className="text-red-600 hover:text-red-900">
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  </div>
);

export default DataTable;