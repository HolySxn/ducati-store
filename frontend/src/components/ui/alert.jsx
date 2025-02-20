import React from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';

const Alert = ({ variant, className, children }) => {
  const alertClass = classNames(
    'p-4 rounded-md flex items-center',
    {
      'bg-red-100 text-red-700': variant === 'destructive',
      'bg-green-100 text-green-700': variant === 'success',
      'bg-yellow-100 text-yellow-700': variant === 'warning',
      'bg-blue-100 text-blue-700': variant === 'info',
    },
    className
  );

  return <div className={alertClass}>{children}</div>;
};

Alert.propTypes = {
  variant: PropTypes.oneOf(['destructive', 'success', 'warning', 'info']),
  className: PropTypes.string,
  children: PropTypes.node.isRequired,
};

const AlertDescription = ({ children }) => {
  return <div className="ml-2">{children}</div>;
};

AlertDescription.propTypes = {
  children: PropTypes.node.isRequired,
};

export { Alert, AlertDescription };