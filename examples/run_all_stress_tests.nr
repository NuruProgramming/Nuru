// Master stress test runner

andika("Starting Nuru Control Flow Stress Tests...")
andika("=" * 50)

andika("\n1. Running While Loop Stress Tests...")
tumia "while_loop_stress"

andika("\n2. Running For...In Loop Stress Tests...")  
tumia "for_in_stress.nr"

andika("\n3. Running Conditional Statement Stress Tests...")
tumia "conditional_stress"

andika("\n4. Running Combined Control Flow Stress Tests...")
tumia "control_flow_combined_stress

andika("\n" + "=" * 50)
andika("All stress tests completed!")
andika("=" * 50)
