import React, { useState } from "react";

export default function GradientDiv() {
  const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });

  // Mouse move event handler
  const handleMouseMove = (event: React.MouseEvent<HTMLDivElement>) => {
    const { clientX, clientY } = event;
    setMousePosition({ x: clientX, y: clientY });
  };

  return (
    <div
      className="absolute w-8 h-8 bg-yellow drop-shadow-yellow"
      style={{
        left: mousePosition.x - 32,
        top: mousePosition.y - 32,
      }}
      onMouseMove={handleMouseMove}
    />
  );
}
