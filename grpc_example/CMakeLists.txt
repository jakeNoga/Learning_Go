cmake_minimum_required(VERSION 3.10)

# Set the project name and language
project(grpc_helloworld C CXX)

include(envCmake.cmake)

set (Protobuf_PROTOC_EXECUTABLE /usr/local/bin/protoc)
set (GRCP_EXECUTABLE /usr/local/bin/grpc_cpp_plugin)
message(STATUS "Protobuf compiler: ${Protobuf_PROTOC_EXECUTABLE}")

# Specify the C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Define directories
set(PROTO_DIR ${CMAKE_CURRENT_SOURCE_DIR}/protos)
set(CPP_OUTPUT_DIR ${PROTO_DIR}/cpp_generated)
file(MAKE_DIRECTORY ${CPP_OUTPUT_DIR})
set(GO_OUTPUT_DIR ${PROTO_DIR}/go_generated)
file(MAKE_DIRECTORY ${GO_OUTPUT_DIR})
set(C_PLUS_SERVER_DIR ${CMAKE_CURRENT_SOURCE_DIR}/grpc_cpp)

# Collect all .proto files in the directory
file(GLOB PROTO_FILES ${PROTO_DIR}/*.proto)

# Debug cmake with:
# message("Proto dir: ${PROTO_DIR}")

# Collect all c plus files in the server directory
file(GLOB C_PLUS_SERVER_SOURCE ${C_PLUS_SERVER_DIR}/*.cc)

# Collect all generated .cc files
set(cpp_grpc_srcs "")

# Function to generate C++ and Go files for a single .proto file
function(add_protobuf_generation proto_file)
    # Get the base filename without extension
    get_filename_component(base_name ${proto_file} NAME_WE)
    get_filename_component(proto_dir ${proto_file} DIRECTORY)
    get_filename_component(proto_file_name ${proto_file} NAME)

    set(cpp_out_files
        ${CPP_OUTPUT_DIR}/${base_name}.pb.cc 
        ${CPP_OUTPUT_DIR}/${base_name}.pb.h
        ${CPP_OUTPUT_DIR}/${base_name}.grpc.pb.h
        ${CPP_OUTPUT_DIR}/${base_name}.grpc.pb.cc)
    set(go_out_file "${GO_OUTPUT_DIR}/${base_name}_grpc.pb.go"
                    ${GO_OUTPUT_DIR}/${base_name}.go)

    message("go_out_file = ${go_out_file}")

    list(APPEND cpp_grpc_srcs ${cpp_out_files})
    set(cpp_grpc_srcs ${cpp_grpc_srcs} PARENT_SCOPE)


    # Add custom command for generating C++ files from .proto files
    add_custom_command(
        OUTPUT ${cpp_out_files}
        COMMAND ${Protobuf_PROTOC_EXECUTABLE}
        ARGS --cpp_out=${CPP_OUTPUT_DIR}
             --grpc_out=${CPP_OUTPUT_DIR}
             --plugin=protoc-gen-grpc=${GRCP_EXECUTABLE}
             --proto_path=${proto_dir}
             ${proto_file_name}
        DEPENDS ${PROTO_FILES}
        COMMENT "Generating C++ protobuf files for ${proto_file}"
        VERBATIM
    )

    # Add custom command for generating Go files from .proto files
    add_custom_command(
        OUTPUT ${go_out_file}
        COMMAND ${Protobuf_PROTOC_EXECUTABLE}
        ARGS --go_out=${proto_dir}
             --go-grpc_out=${proto_dir}
             --proto_path=${proto_dir}
             ${proto_file}
        DEPENDS ${proto_file}
        COMMENT "Generating Go protobuf files for ${proto_file}"
        VERBATIM
    )

    # Add custom targets for each file
    add_custom_target(
        generate_cpp_${base_name}
        ALL
        DEPENDS ${cpp_out_files}
    )

    # Add custom target for each file
    add_custom_target(
        generate_go_${base_name}
        ALL
        DEPENDS ${go_out_file}
    )
endfunction(add_protobuf_generation proto_file)


# message("Proto files: ${PROTO_FILES}")
# Apply the function to all .proto files
foreach(proto_file ${PROTO_FILES})
    message("Generating C++ protobuf files for ${proto_file}")
    add_protobuf_generation(${proto_file})
endforeach(proto_file ${PROTO_FILES})

# Define the hw_grpc_proto library
add_library(hw_grpc_proto ${cpp_grpc_srcs})
target_link_libraries(hw_grpc_proto
  absl::check
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF})

# Define targets for each server file
foreach(_source_file ${C_PLUS_SERVER_SOURCE})
  get_filename_component(_target_name ${_source_file} NAME_WE)

  add_executable(${_target_name} "${_source_file}")

  target_link_libraries(${_target_name}
    hw_grpc_proto
    absl::check
    absl::flags
    absl::flags_parse
    absl::log
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF})

endforeach()