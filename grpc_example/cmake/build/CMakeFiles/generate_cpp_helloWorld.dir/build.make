# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.19

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Disable VCS-based implicit rules.
% : %,v


# Disable VCS-based implicit rules.
% : RCS/%


# Disable VCS-based implicit rules.
% : RCS/%,v


# Disable VCS-based implicit rules.
% : SCCS/s.%


# Disable VCS-based implicit rules.
% : s.%


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /home/noga/.local/bin/cmake

# The command to remove a file.
RM = /home/noga/.local/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /mnt/c/Users/nogaj/Documents/learning_go/grpc_example

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /mnt/c/Users/nogaj/Documents/learning_go/grpc_example/cmake/build

# Utility rule file for generate_cpp_helloWorld.

# Include the progress variables for this target.
include CMakeFiles/generate_cpp_helloWorld.dir/progress.make

CMakeFiles/generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.pb.cc
CMakeFiles/generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.pb.h
CMakeFiles/generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.grpc.pb.h
CMakeFiles/generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.grpc.pb.cc


../../protos/cpp_generated/helloWorld.pb.cc: ../../protos/helloWorld.proto
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --blue --bold --progress-dir=/mnt/c/Users/nogaj/Documents/learning_go/grpc_example/cmake/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Generating C++ protobuf files for /mnt/c/Users/nogaj/Documents/learning_go/grpc_example/protos/helloWorld.proto"
	/usr/local/bin/protoc --cpp_out=/mnt/c/Users/nogaj/Documents/learning_go/grpc_example/protos/cpp_generated --grpc_out=/mnt/c/Users/nogaj/Documents/learning_go/grpc_example/protos/cpp_generated --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin --proto_path=/mnt/c/Users/nogaj/Documents/learning_go/grpc_example/protos helloWorld.proto

../../protos/cpp_generated/helloWorld.pb.h: ../../protos/cpp_generated/helloWorld.pb.cc
	@$(CMAKE_COMMAND) -E touch_nocreate ../../protos/cpp_generated/helloWorld.pb.h

../../protos/cpp_generated/helloWorld.grpc.pb.h: ../../protos/cpp_generated/helloWorld.pb.cc
	@$(CMAKE_COMMAND) -E touch_nocreate ../../protos/cpp_generated/helloWorld.grpc.pb.h

../../protos/cpp_generated/helloWorld.grpc.pb.cc: ../../protos/cpp_generated/helloWorld.pb.cc
	@$(CMAKE_COMMAND) -E touch_nocreate ../../protos/cpp_generated/helloWorld.grpc.pb.cc

generate_cpp_helloWorld: CMakeFiles/generate_cpp_helloWorld
generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.grpc.pb.cc
generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.grpc.pb.h
generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.pb.cc
generate_cpp_helloWorld: ../../protos/cpp_generated/helloWorld.pb.h
generate_cpp_helloWorld: CMakeFiles/generate_cpp_helloWorld.dir/build.make

.PHONY : generate_cpp_helloWorld

# Rule to build all files generated by this target.
CMakeFiles/generate_cpp_helloWorld.dir/build: generate_cpp_helloWorld

.PHONY : CMakeFiles/generate_cpp_helloWorld.dir/build

CMakeFiles/generate_cpp_helloWorld.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/generate_cpp_helloWorld.dir/cmake_clean.cmake
.PHONY : CMakeFiles/generate_cpp_helloWorld.dir/clean

CMakeFiles/generate_cpp_helloWorld.dir/depend:
	cd /mnt/c/Users/nogaj/Documents/learning_go/grpc_example/cmake/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /mnt/c/Users/nogaj/Documents/learning_go/grpc_example /mnt/c/Users/nogaj/Documents/learning_go/grpc_example /mnt/c/Users/nogaj/Documents/learning_go/grpc_example/cmake/build /mnt/c/Users/nogaj/Documents/learning_go/grpc_example/cmake/build /mnt/c/Users/nogaj/Documents/learning_go/grpc_example/cmake/build/CMakeFiles/generate_cpp_helloWorld.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/generate_cpp_helloWorld.dir/depend

